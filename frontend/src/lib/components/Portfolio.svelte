<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
  const projects = [
    {
      title: 'سامانه مدیریت پروژه هوشمند',
      stack: ['SvelteKit', 'Go Fiber', 'PostgreSQL', 'Docker'],
      result: 'افزایش بهره‌وری تیم ۳۵٪',
      link: '#'
    },
    {
      title: 'بات تلگرام فروشگاهی',
      stack: ['Node.js', 'RAG', 'Redis'],
      result: 'افزایش فروش ۲.۲ برابر',
      link: '#'
    },
    {
      title: 'داشبورد تحلیل داده معاملاتی',
      stack: ['Python', 'FastAPI', 'React', 'Kafka'],
      result: 'کاهش زمان تحلیل ۵۰٪',
      link: '#'
    }
  ];
</script>

<section id="portfolio" class="section relative">
  <div class="container-rtl">
    <div class="text-center max-w-2xl mx-auto" use:reveal>
      <span class="eyebrow">نمونه کارها</span>
      <h2 class="mt-3 text-3xl md:text-4xl font-bold tracking-tight">مروری بر پروژه‌های شاخص</h2>
      <p class="lead mt-2 text-slate-600 dark:text-slate-300">منتخبی از دستاوردهای اخیر ما</p>
    </div>

    <div class="mt-8 grid gap-6 md:grid-cols-2 lg:grid-cols-3" role="list" aria-label="پروژه‌ها">
      {#each projects as p, i}
        <article class="group card p-0 overflow-hidden flex flex-col rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm transition-all hover:shadow-lg hover:-translate-y-1 motion-reduce:hover:translate-y-0 relative focus-within:ring-2 focus-within:ring-primary-500/60 focus-within:ring-offset-2 focus-within:ring-offset-transparent" use:reveal role="listitem" aria-labelledby={'proj-' + i}>
          <!-- Decorative header strip -->
          <div class="h-1.5 w-full bg-gradient-to-l from-primary-500/60 via-cyan-400/60 to-primary-500/60"></div>
          <div class="p-6">
            <h3 class="text-xl font-semibold" id={'proj-' + i}>{p.title}</h3>
            <div class="mt-3 flex flex-wrap gap-2 text-xs text-slate-600 dark:text-slate-300">
              {#each p.stack as s}
                <span class="badge">{s}</span>
              {/each}
            </div>
            <p class="mt-4 text-sm text-slate-600 dark:text-slate-300">نتیجه: {p.result}</p>
          </div>
          <div class="mt-auto p-6 pt-0">
            <a href={p.link} class="inline-flex items-center gap-1 text-primary-700 dark:text-primary-300 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 rounded transition-colors hover:text-primary-800 dark:hover:text-primary-200" aria-label={`مشاهده جزئیات پروژه ${p.title}`}>
              مشاهده جزئیات
              <svg class="w-4 h-4 transition-transform duration-200 group-hover:translate-x-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M5 12h14"/><path d="M12 5l7 7-7 7"/></svg>
            </a>
          </div>
        </article>
      {/each}
    </div>
  </div>
</section>
