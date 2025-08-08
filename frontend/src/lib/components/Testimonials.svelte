<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
  const items = [
    { name: 'شرکت آلفا', text: 'کیفیت توسعه عالی و تحویل به‌موقع. همکاری با تهران‌بات بسیار حرفه‌ای بود.', rating: 5 },
    { name: 'استارتاپ بتا', text: 'ربات هوشمند طراحی‌شده عملکرد بی‌نظیری دارد و باعث افزایش رضایت مشتریان شد.', rating: 5 },
    { name: 'هلدینگ گاما', text: 'تیم فنی مسلط و پاسخگو. نتایج فراتر از انتظار بود.', rating: 4 }
  ];
</script>

<section class="section relative">
  <div class="container-rtl">
    <div class="text-center max-w-2xl mx-auto" use:reveal>
      <span class="inline-flex items-center gap-2 rounded-full border border-primary-500/20 bg-primary-500/10 px-3 py-1 text-sm text-primary-700 dark:text-primary-300">نظرات مشتریان</span>
      <h2 class="mt-3 text-3xl md:text-4xl font-bold tracking-tight">صدای مشتریان ما</h2>
      <p class="lead mt-2 text-slate-600 dark:text-slate-300">اعتماد کسب‌وکارها به نتایج ما</p>
    </div>

    <div class="mt-8 grid gap-6 md:grid-cols-3">
      {#each items as t}
        <figure class="card p-6 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm" use:reveal>
          <figcaption class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-9 h-9 rounded-full bg-primary-500/15 text-primary-700 dark:text-primary-300 grid place-items-center font-bold">{t.name?.[0]}</div>
              <span class="font-semibold">{t.name}</span>
            </div>
            <div class="flex text-amber-500" aria-label={`امتیاز ${t.rating} از ۵`}>
              {#each Array(5) as _, i}
                <svg class="w-5 h-5" aria-hidden="true" fill={i < t.rating ? 'currentColor' : 'none'} stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557L3.04 10.385a.563.562 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345l2.125-5.111z"/></svg>
              {/each}
            </div>
          </figcaption>
          <span class="sr-only">امتیاز: {t.rating} از ۵</span>
          <blockquote class="mt-3 text-slate-700 dark:text-slate-300">{t.text}</blockquote>
        </figure>
      {/each}
    </div>
  </div>
</section>
