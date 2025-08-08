<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
  const services = [
    {
      title: 'توسعه و معماری نرم‌افزار',
      icon: 'layers',
      items: [
        'طراحی معماری نرم‌افزار و پایگاه داده',
        'توسعه نرم‌افزارهای تحت وب',
        'توسعه اپلیکیشن‌های موبایل',
        'طراحی و توسعه وب‌سایت'
      ]
    },
    {
      title: 'هوش مصنوعی و اتوماسیون',
      icon: 'spark',
      items: [
        'ساخت ربات‌های هوشمند AI',
        'سیستم‌های RAG',
        'ربات‌های تلگرام، اینستاگرام، واتساپ'
      ]
    },
    {
      title: 'خدمات تخصصی',
      icon: 'shield',
      items: [
        'اپلیکیشن‌های معاملاتی',
        'بهینه‌سازی موتورهای جستجو (SEO)',
        'DevOps و MLOps',
        'اتوماسیون فرآیندهای کسب‌وکار'
      ]
    }
  ];
</script>

<section id="services" class="section relative">
  <!-- Decorative background -->
  <div aria-hidden="true" class="pointer-events-none absolute inset-x-0 -top-8 h-40 bg-gradient-to-b from-primary-500/10 to-transparent blur-2xl"></div>

  <div class="container-rtl">
    <div class="text-center max-w-2xl mx-auto" use:reveal>
      <span class="inline-flex items-center gap-2 rounded-full border border-primary-500/20 bg-primary-500/10 px-3 py-1 text-sm text-primary-700 dark:text-primary-300">
        <span class="i-lucide-stars" aria-hidden="true"></span>
        آنچه ارائه می‌دهیم
      </span>
      <h2 class="mt-3 text-3xl md:text-4xl font-bold tracking-tight">خدمات ما</h2>
      <p class="lead mt-2 text-slate-600 dark:text-slate-300">راهکارهای به‌روز بر پایه تکنولوژی‌های مدرن، متناسب با نیاز کسب‌وکار شما</p>
    </div>

    <div class="mt-8 grid gap-6 sm:grid-cols-2 lg:grid-cols-3" role="list">
      {#each services as s, i}
        <article
          class="card group p-6 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm hover:shadow-lg hover:-translate-y-1 transition-all focus-within:ring-2 focus-within:ring-primary-500"
          aria-labelledby={`service-${i}`}
          use:reveal
          role="listitem"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-xl bg-gradient-to-br from-primary-500/15 to-primary-500/5 text-primary-600 dark:text-primary-400 grid place-items-center">
                {#if s.icon === 'layers'}
                  <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M12 2l9 4-9 4-9-4 9-4z"/><path d="M3 10l9 4 9-4"/><path d="M3 18l9 4 9-4"/></svg>
                {:else if s.icon === 'spark'}
                  <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M5 3l4 9-4 9 9-4 9 4-4-9 4-9-9 4-9-4z"/></svg>
                {:else if s.icon === 'shield'}
                  <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/></svg>
                {/if}
              </div>
              <h3 id={`service-${i}`} class="text-xl font-semibold">{s.title}</h3>
            </div>
          </div>

          <ul class="mt-4 space-y-2 text-slate-600 dark:text-slate-300">
            {#each s.items as it}
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 mt-0.5 text-primary-600 shrink-0" viewBox="0 0 20 20" fill="currentColor" xmlns="http://www.w3.org/2000/svg" aria-hidden="true"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/></svg>
                <span>{it}</span>
              </li>
            {/each}
          </ul>

          <div class="mt-5 flex items-center justify-end">
            <a href="#contact" class="text-primary-700 dark:text-primary-300 hover:underline inline-flex items-center gap-1">
              مشاهده جزئیات
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M5 12h14"/><path d="M12 5l7 7-7 7"/></svg>
            </a>
          </div>
        </article>
      {/each}
    </div>
  </div>
</section>
