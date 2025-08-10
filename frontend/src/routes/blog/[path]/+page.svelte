<script lang="ts">
  export let data: { blog: { category: string; text: string; path: string }; similar: { category: string; text: string; path: string }[] };
  import { onMount } from 'svelte';
  import { env as publicEnv } from '$env/dynamic/public';

  function stripHtml(html: string): string {
    return html.replace(/<[^>]*>/g, ' ');
  }

  function readingTime(html: string): number {
    const words = stripHtml(html).trim().split(/\s+/).filter(Boolean).length;
    return Math.max(1, Math.round(words / 200));
  }

  const faDigits = ['۰','۱','۲','۳','۴','۵','۶','۷','۸','۹'];
  function faNum(n: number): string {
    return String(n).replace(/\d/g, (d) => faDigits[Number(d)] ?? d);
  }

  // Derive meta and display values from HTML body
  function titleFromHTML(html: string, fallback: string): string {
    const h1 = /<h1[^>]*>([\s\S]*?)<\/h1>/i.exec(html)?.[1];
    if (h1) return stripHtml(h1).trim();
    const h2 = /<h2[^>]*>([\s\S]*?)<\/h2>/i.exec(html)?.[1];
    if (h2) return stripHtml(h2).trim();
    return fallback;
  }
  function descriptionFromHTML(html: string, max = 160): string {
    const p = /<p[^>]*>([\s\S]*?)<\/p>/i.exec(html)?.[1];
    const text = stripHtml(p || html).replace(/\s+/g, ' ').trim();
    return text.length > max ? (text.slice(0, max).trim() + '…') : text;
  }
  function firstImage(html: string): { src: string; alt: string } | null {
    const imgMatch = /<img[^>]*src=["']([^"']+)["'][^>]*>/i.exec(html);
    if (!imgMatch) return null;
    const tag = imgMatch[0];
    const src = imgMatch[1];
    const altMatch = /alt=["']([^"']*)["']/i.exec(tag);
    const alt = altMatch ? altMatch[1] : '';
    return { src, alt };
  }

  // Site/Canonical
  const BASE = (publicEnv.PUBLIC_BASE_URL?.trim()) || 'https://tehranbot.me';
  const SITE_NAME = 'TehranBot';
  $: canonical = `${BASE}/blog/${encodeURIComponent(data.blog.path)}`;
  $: pageTitle = `${titleFromHTML(data.blog.text, data.blog.path)} | ${SITE_NAME}`;
  $: metaDescription = descriptionFromHTML(data.blog.text, 180);

  // Reading progress
  let progress = 0;
  let articleEl: HTMLElement;
  let articleTop = 0;
  let articleHeight = 0;
  let activeId: string | null = null;

  function slugify(text: string): string {
    return text
      .trim()
      .toLowerCase()
      .replace(/[^\p{L}\p{N}\s-]/gu, '')
      .replace(/\s+/g, '-')
      .replace(/-+/g, '-');
  }

  type TocItem = { id: string; text: string; level: number };
  let toc: TocItem[] = [];

  function recalc() {
    if (!articleEl) return;
    const rect = articleEl.getBoundingClientRect();
    articleTop = rect.top + window.scrollY;
    articleHeight = articleEl.offsetHeight;
    updateProgress();
  }

  function updateProgress() {
    if (!articleEl) return;
    const start = articleTop;
    const end = Math.max(start, articleTop + articleHeight - window.innerHeight);
    const y = window.scrollY;
    const denom = end - start || 1;
    const p = ((y - start) / denom) * 100;
    progress = Math.max(0, Math.min(100, p));
  }

  onMount(() => {
    // Build ToC from h2/h3
    if (articleEl) {
      const headers = Array.from(articleEl.querySelectorAll('h2, h3')) as HTMLElement[];
      const seen = new Set<string>();
      toc = headers.map((h) => {
        let id = h.id || slugify(h.textContent || '');
        let base = id || 'section';
        let unique = base;
        let i = 1;
        while (seen.has(unique)) {
          unique = `${base}-${i++}`;
        }
        seen.add(unique);
        if (!h.id) h.id = unique;
        return { id: unique, text: (h.textContent || '').trim(), level: h.tagName === 'H3' ? 3 : 2 };
      });

      // Active section tracking via IntersectionObserver
      const io = new IntersectionObserver(
        (entries) => {
          // Pick the heading nearest to top that's intersecting
          const visible = entries
            .filter((e) => e.isIntersecting)
            .sort((a, b) => (a.target as HTMLElement).offsetTop - (b.target as HTMLElement).offsetTop);
          if (visible.length > 0) {
            activeId = (visible[0].target as HTMLElement).id || null;
          }
        },
        {
          root: null,
          // Trigger a bit before the top to feel responsive
          rootMargin: '-20% 0px -70% 0px',
          threshold: [0, 1]
        }
      );
      headers.forEach((h) => io.observe(h));
    }

    recalc();
    const onScroll = () => updateProgress();
    const onResize = () => recalc();
    window.addEventListener('scroll', onScroll, { passive: true });
    window.addEventListener('resize', onResize);
    const ro = new ResizeObserver(() => recalc());
    if (articleEl) ro.observe(articleEl);
    return () => {
      window.removeEventListener('scroll', onScroll as any);
      window.removeEventListener('resize', onResize as any);
      ro.disconnect();
    };
  });
</script>

<svelte:head>
  <link rel="canonical" href={canonical} />
  <title>{pageTitle}</title>
  <meta name="description" content={metaDescription} />
  <meta property="og:type" content="article" />
  <meta property="og:title" content={pageTitle} />
  <meta property="og:description" content={metaDescription} />
  <meta property="og:url" content={canonical} />
  {#if firstImage(data.blog.text)}
    <meta property="og:image" content={firstImage(data.blog.text)?.src} />
  {/if}
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:title" content={pageTitle} />
  <meta name="twitter:description" content={metaDescription} />
  {#if firstImage(data.blog.text)}
    <meta name="twitter:image" content={firstImage(data.blog.text)?.src} />
  {/if}
</svelte:head>

<section id="main-content" class="container-rtl py-10">
  <!-- Top reading progress bar -->
  <div class="fixed top-0 inset-x-0 h-1 bg-transparent z-40">
    <div class="h-full bg-primary-500/80" style={`width:${progress}%; transition: width 100ms linear;`}></div>
  </div>

  <div class="mx-auto max-w-6xl lg:grid lg:grid-cols-[280px_minmax(0,1fr)] lg:gap-8">
    <div class="min-w-0 lg:col-start-2 lg:row-start-1">
      <nav class="mb-3 text-sm text-slate-500">
        <a class="hover:underline" href="/blog">بلاگ</a>
        <span class="mx-1">/</span>
        <span class="text-slate-700 dark:text-slate-300">{titleFromHTML(data.blog.text, data.blog.path)}</span>
      </nav>

      <div class="mb-2 flex items-center gap-2">
        <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-300 text-[11px]">{data.blog.category}</span>
        <span class="text-[11px] text-slate-400">{faNum(readingTime(data.blog.text))} دقیقه مطالعه</span>
      </div>
      <h1 class="text-3xl font-extrabold tracking-tight mb-6">{titleFromHTML(data.blog.text, data.blog.path)}</h1>

      {#if toc.length > 0}
        <!-- Mobile/Tablet ToC -->
        <aside class="mb-6 border rounded-xl bg-white/70 dark:bg-slate-900/40 p-4 lg:hidden">
          <h2 class="text-sm font-bold mb-2 text-slate-700 dark:text-slate-200">فهرست مطالب</h2>
          <nav>
            <ul class="space-y-1 text-sm">
              {#each toc as item}
                <li class={`ps-${item.level === 3 ? '4' : '0'}`}>
                  <a class={`hover:underline ${activeId === item.id ? 'text-primary-600 dark:text-primary-400' : 'text-slate-600 dark:text-slate-300'}`} href={`#${item.id}`} aria-current={activeId === item.id ? 'true' : 'false'}>{item.text}</a>
                </li>
              {/each}
            </ul>
          </nav>
        </aside>
      {/if}

    <article bind:this={articleEl} class="prose prose-lg md:prose-xl prose-slate dark:prose-invert max-w-none leading-relaxed prose-a:no-underline hover:prose-a:underline prose-img:rounded-xl prose-blockquote:border-s-4 prose-blockquote:ps-4 prose-pre:rounded-xl prose-pre:bg-slate-900 prose-pre:text-slate-100 prose-code:px-1 prose-code:py-0.5 prose-code:bg-slate-100 dark:prose-code:bg-slate-800 prose-headings:scroll-mt-24">
      {@html data.blog.text}
    </article>

    {#if data.similar && data.similar.length > 0}
      <div class="mt-10 border-t pt-6">
        <h2 class="text-lg font-bold mb-3 text-slate-800 dark:text-slate-200">مقالات مشابه</h2>
        <ul class="grid gap-3 sm:grid-cols-2">
          {#each data.similar as s}
            <li class="group border rounded-xl p-4 bg-white/70 dark:bg-slate-900/40 hover:border-primary-400 transition-colors">
              <a class="block" href={`/blog/${encodeURIComponent(s.path)}`}>
                <div class="mb-2 flex items-center gap-2">
                  <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-300 text-[11px]">{s.category}</span>
                </div>
                <h3 class="text-sm font-bold text-slate-800 dark:text-slate-200 line-clamp-1">{s.path}</h3>
                <p class="mt-1 text-xs text-slate-500 dark:text-slate-400 line-clamp-2">{stripHtml(s.text).slice(0, 120)}{stripHtml(s.text).length > 120 ? '…' : ''}</p>
              </a>
            </li>
          {/each}
        </ul>
      </div>
    {/if}
    </div>

    {#if toc.length > 0}
      <!-- Desktop sticky ToC on the right -->
      <aside class="hidden lg:block lg:col-start-1 lg:row-start-1 self-start">
        <div class="sticky top-24 max-h-[calc(100vh-6rem)] overflow-auto border rounded-xl bg-white/70 dark:bg-slate-900/40 p-4">
          <h2 class="text-sm font-bold mb-2 text-slate-700 dark:text-slate-200">فهرست مطالب</h2>
          <nav>
            <ul class="space-y-1 text-sm">
              {#each toc as item}
                <li class={`ps-${item.level === 3 ? '4' : '0'}`}>
                  <a class={`hover:underline ${activeId === item.id ? 'text-primary-600 dark:text-primary-400' : 'text-slate-600 dark:text-slate-300'}`} href={`#${item.id}`} aria-current={activeId === item.id ? 'true' : 'false'}>{item.text}</a>
                </li>
              {/each}
            </ul>
          </nav>
        </div>
      </aside>
    {/if}
  </div>
</section>
