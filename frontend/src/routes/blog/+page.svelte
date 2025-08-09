<script lang="ts">
  export let data: { blogs: { category: string; text: string; path: string }[]; category: string };
  import { goto } from '$app/navigation';

  const categories = Array.from(new Set(data.blogs.map((b) => b.category)));

  function setCategory(cat: string | null) {
    const params = new URLSearchParams(typeof window !== 'undefined' ? window.location.search : '');
    if (cat && cat.trim()) params.set('category', cat);
    else params.delete('category');
    goto(`/blog${params.toString() ? `?${params.toString()}` : ''}`, { replaceState: true });
  }
</script>

<section class="container-rtl py-10">
  <h1 class="text-2xl font-extrabold mb-6">بلاگ</h1>

  <div class="flex flex-wrap gap-2 mb-6">
    <button class="px-3 py-1 rounded-full border text-sm hover:bg-slate-100 dark:hover:bg-slate-800"
      class:selected={!data.category}
      on:click={() => setCategory(null)}>
      همه
    </button>
    {#each categories as c}
      <button class="px-3 py-1 rounded-full border text-sm hover:bg-slate-100 dark:hover:bg-slate-800"
        class:selected={data.category === c}
        on:click={() => setCategory(c)}>
        {c}
      </button>
    {/each}
  </div>

  {#if data.blogs.length === 0}
    <div class="text-slate-500">هیچ مقاله‌ای یافت نشد.</div>
  {:else}
    <ul class="grid grid-cols-1 md:grid-cols-2 gap-6">
      {#each data.blogs as b}
        <li class="rounded-xl border p-4 bg-white/60 dark:bg-slate-900/40">
          <a class="block" href={`/blog/${b.path}`}>
            <div class="text-xs text-slate-500 mb-1">{b.category}</div>
            <div class="font-bold mb-2">{b.path}</div>
            <div class="prose prose-slate dark:prose-invert line-clamp-3" {@html b.text}></div>
          </a>
        </li>
      {/each}
    </ul>
  {/if}
</section>
