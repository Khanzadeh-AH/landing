<script lang="ts">
  export let data: { blog: { category: string; text: string; path: string }; similar: { category: string; text: string; path: string }[] };
  import { onMount } from 'svelte';

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

  // Reading progress
  let progress = 0;
  let articleEl: HTMLElement;
  let articleTop = 0;
  let articleHeight = 0;

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

<section class="container-rtl py-10">
  <!-- Top reading progress bar -->
  <div class="fixed top-0 inset-x-0 h-1 bg-transparent z-40">
    <div class="h-full bg-primary-500/80" style={`width:${progress}%; transition: width 100ms linear;`}></div>
  </div>

  <div class="max-w-3xl mx-auto">
    <nav class="mb-3 text-sm text-slate-500">
      <a class="hover:underline" href="/blog">بلاگ</a>
      <span class="mx-1">/</span>
      <span class="text-slate-700 dark:text-slate-300">{data.blog.path}</span>
    </nav>

    <div class="mb-2 flex items-center gap-2">
      <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-300 text-[11px]">{data.blog.category}</span>
      <span class="text-[11px] text-slate-400">{faNum(readingTime(data.blog.text))} دقیقه مطالعه</span>
    </div>

    <h1 class="text-3xl font-extrabold tracking-tight mb-6">{data.blog.path}</h1>

    {#if toc.length > 0}
      <aside class="mb-6 border rounded-xl bg-white/70 dark:bg-slate-900/40 p-4">
        <h2 class="text-sm font-bold mb-2 text-slate-700 dark:text-slate-200">فهرست مطالب</h2>
        <nav>
          <ul class="space-y-1 text-sm">
            {#each toc as item}
              <li class={`ps-${item.level === 3 ? '4' : '0'}`}>
                <a class="hover:underline text-slate-600 dark:text-slate-300" href={`#${item.id}`}>{item.text}</a>
              </li>
            {/each}
          </ul>
        </nav>
      </aside>
    {/if}

    <article bind:this={articleEl} class="prose prose-slate dark:prose-invert max-w-none">
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
</section>
