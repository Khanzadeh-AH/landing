<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
  const services = [
    {
      title: 'توسعه و معماری نرم‌افزار',
      items: [
        'طراحی معماری نرم‌افزار و پایگاه داده',
        'توسعه نرم‌افزارهای تحت وب',
        'توسعه اپلیکیشن‌های موبایل',
        'طراحی و توسعه وب‌سایت'
      ]
    },
    {
      title: 'هوش مصنوعی و اتوماسیون',
      items: [
        'ساخت ربات‌های هوشمند AI',
        'سیستم‌های RAG',
        'ربات‌های تلگرام، اینستاگرام، واتساپ'
      ]
    },
    {
      title: 'خدمات تخصصی',
      items: [
        'اپلیکیشن‌های معاملاتی',
        'بهینه‌سازی موتورهای جستجو (SEO)',
        'DevOps و MLOps',
        'اتوماسیون فرآیندهای کسب‌وکار'
      ]
    }
  ];
</script>

<section id="services" class="section">
  <div class="container-rtl">
    <div class="text-center" use:reveal>
      <h2 class="text-3xl md:text-4xl">خدمات ما</h2>
      <p class="lead mt-2">راهکارهای به‌روز بر پایه تکنولوژی‌های مدرن</p>
    </div>

    <div class="mt-8 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {#each services as s}
        <div class="card p-6" use:reveal>
          <h3 class="text-xl">{s.title}</h3>
          <ul class="mt-4 space-y-2 text-slate-600 dark:text-slate-300">
            {#each s.items as it}
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-primary-600 shrink-0" viewBox="0 0 20 20" fill="currentColor" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/></svg>
                <span>{it}</span>
              </li>
            {/each}
          </ul>
        </div>
      {/each}
    </div>
  </div>
</section>
