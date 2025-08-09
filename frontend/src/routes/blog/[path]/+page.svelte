<script lang="ts">
  export let data: { blog: { category: string; text: string; path: string } };

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
</script>

<section class="container-rtl py-10">
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

    <article class="prose prose-slate dark:prose-invert max-w-none">
      {@html data.blog.text}
    </article>
  </div>
</section>
