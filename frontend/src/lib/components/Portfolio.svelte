<script lang="ts">
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

<section id="portfolio" class="section">
  <div class="container-rtl">
    <div class="text-center">
      <h2 class="text-3xl md:text-4xl">نمونه کارها</h2>
      <p class="lead mt-2">مروری بر برخی پروژه‌های شاخص</p>
    </div>

    <div class="mt-8 grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      {#each projects as p}
        <article class="card p-6 flex flex-col">
          <h3 class="text-xl">{p.title}</h3>
          <div class="mt-3 flex flex-wrap gap-2 text-xs text-slate-600 dark:text-slate-300">
            {#each p.stack as s}
              <span class="badge">{s}</span>
            {/each}
          </div>
          <p class="mt-4 text-sm text-slate-600 dark:text-slate-300">نتیجه: {p.result}</p>
          <div class="mt-auto pt-4">
            <a href={p.link} class="text-primary-600 hover:underline">مشاهده جزئیات</a>
          </div>
        </article>
      {/each}
    </div>
  </div>
</section>
