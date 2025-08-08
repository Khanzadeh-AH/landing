<script lang="ts">
  const items = [
    { name: 'شرکت آلفا', text: 'کیفیت توسعه عالی و تحویل به‌موقع. همکاری با تهران‌بات بسیار حرفه‌ای بود.', rating: 5 },
    { name: 'استارتاپ بتا', text: 'ربات هوشمند طراحی‌شده عملکرد بی‌نظیری دارد و باعث افزایش رضایت مشتریان شد.', rating: 5 },
    { name: 'هلدینگ گاما', text: 'تیم فنی مسلط و پاسخگو. نتایج فراتر از انتظار بود.', rating: 4 }
  ];
</script>

<section class="section">
  <div class="container-rtl">
    <div class="text-center">
      <h2 class="text-3xl md:text-4xl">نظرات مشتریان</h2>
      <p class="lead mt-2">اعتماد کسب‌وکارها به نتایج ما</p>
    </div>

    <div class="mt-8 grid gap-6 md:grid-cols-3">
      {#each items as t}
        <div class="card p-6">
          <div class="flex items-center justify-between">
            <div class="font-semibold">{t.name}</div>
            <div class="flex text-yellow-500">
              {#each Array(5) as _, i}
                <svg class="w-5 h-5" fill={i < t.rating ? 'currentColor' : 'none'} stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557L3.04 10.385a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345l2.125-5.111z"/></svg>
              {/each}
            </div>
          </div>
          <p class="mt-3 text-slate-700 dark:text-slate-300">{t.text}</p>
        </div>
      {/each}
    </div>
  </div>
</section>
