<script lang="ts">
  export let data: { blog: { category: string; text: string; path: string } };
</script>

<section class="container-rtl py-10">
  <div class="text-xs text-slate-500 mb-2">{data.blog.category}</div>
  <h1 class="text-2xl font-extrabold mb-6">{data.blog.path}</h1>
  <article class="prose prose-slate dark:prose-invert max-w-none" {@html data.blog.text}></article>
</section>
