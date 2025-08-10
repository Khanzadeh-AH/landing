<script lang="ts">
  export let data: { blogs: { category: string; text: string; path: string }[]; category: string };
  import { goto } from '$app/navigation';

  // Categories are derived from the current dataset (may be filtered server-side by category)
  const categories = Array.from(new Set(data.blogs.map((b) => b.category)));

  // Local UI state
  let query = '';
  let visibleCount = 10;

  // Helpers
  function stripHtml(html: string): string {
    return html.replace(/<[^>]*>/g, ' ');
  }

  function excerpt(html: string, max = 160): string {
    const text = stripHtml(html).replace(/\s+/g, ' ').trim();
    if (text.length <= max) return text;
    return text.slice(0, max).trim() + '…';
  }

  function readingTime(html: string): number {
    const words = stripHtml(html).trim().split(/\s+/).filter(Boolean).length;
    return Math.max(1, Math.round(words / 200)); // ~200 wpm
  }

  const faDigits = ['۰','۱','۲','۳','۴','۵','۶','۷','۸','۹'];
  function faNum(n: number): string {
    return String(n).replace(/\d/g, (d) => faDigits[Number(d)] ?? d);
  }

  function setCategory(cat: string | null) {
    const params = new URLSearchParams(typeof window !== 'undefined' ? window.location.search : '');
    if (cat && cat.trim()) params.set('category', cat);
    else params.delete('category');
    goto(`/blog${params.toString() ? `?${params.toString()}` : ''}`, { replaceState: true });
  }

  // Client-side search and pagination over the already-fetched list
  $: filtered = data.blogs.filter((b) => {
    if (!query.trim()) return true;
    const haystack = (b.path + ' ' + stripHtml(b.text)).toLowerCase();
    return haystack.includes(query.trim().toLowerCase());
  });

  $: visible = filtered.slice(0, visibleCount);

  function loadMore() {
    visibleCount += 10;
  }
</script>

<section id="main-content" class="container-rtl py-10">
  <header class="mb-6">
    <h1 class="text-3xl font-extrabold tracking-tight mb-2">بلاگ</h1>
    <p class="text-slate-500 text-sm">مطالب جدید و به‌روز را در اینجا بخوانید.</p>
  </header>

  <div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between mb-6">
    <div class="flex flex-wrap gap-2">
      <button
        class={`px-3 py-1 rounded-full border text-sm transition ${!data.category ? 'bg-slate-900 text-white dark:bg-slate-100 dark:text-slate-900' : 'hover:bg-slate-100 dark:hover:bg-slate-800'}`}
        aria-pressed={!data.category}
        on:click={() => setCategory(null)}
      >
        همه
      </button>
      {#each categories as c}
        <button
          class={`px-3 py-1 rounded-full border text-sm transition ${data.category === c ? 'bg-slate-900 text-white dark:bg-slate-100 dark:text-slate-900' : 'hover:bg-slate-100 dark:hover:bg-slate-800'}`}
          aria-pressed={data.category === c}
          on:click={() => setCategory(c)}
        >
          {c}
        </button>
      {/each}
    </div>
    <div class="w-full md:w-72">
      <label class="sr-only" for="blog-search">جستجو</label>
      <input
        id="blog-search"
        class="w-full px-3 py-2 rounded-lg border bg-white/70 dark:bg-slate-900/50 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-slate-400 dark:focus:ring-slate-600"
        type="search"
        placeholder="جستجو در مقالات..."
        bind:value={query}
      />
    </div>
  </div>

  {#if filtered.length === 0}
    <div class="text-slate-500 bg-slate-50 dark:bg-slate-900/40 border rounded-xl p-6">
      هیچ مقاله‌ای با «{query}» یافت نشد.
    </div>
  {:else}
    <ul class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-6">
      {#each visible as b}
        <li class="group rounded-2xl border p-5 bg-white/70 dark:bg-slate-900/40 transition hover:shadow-md hover:-translate-y-0.5">
          <a class="block focus:outline-none focus:ring-2 focus:ring-slate-400 dark:focus:ring-slate-600 rounded-xl" href={`/blog/${b.path}`}>
            <div class="mb-2 flex items-center gap-2">
              <span class="inline-flex items-center px-2 py-0.5 rounded-full bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-300 text-[11px]">{b.category}</span>
              <span class="text-[11px] text-slate-400">{faNum(readingTime(b.text))} دقیقه مطالعه</span>
            </div>
            <h2 class="font-extrabold text-base mb-2 line-clamp-2 group-hover:underline">{b.path}</h2>
            <p class="text-sm text-slate-600 dark:text-slate-300 line-clamp-3">{excerpt(b.text, 180)}</p>
          </a>
        </li>
      {/each}
    </ul>

    {#if filtered.length > visibleCount}
      <div class="flex justify-center mt-8">
        <button
          class="px-4 py-2 rounded-lg border bg-white/70 dark:bg-slate-900/40 hover:bg-slate-100 dark:hover:bg-slate-800 transition"
          on:click={loadMore}
        >
          نمایش بیشتر
        </button>
      </div>
    {/if}
  {/if}
</section>
