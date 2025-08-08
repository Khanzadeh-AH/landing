<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
</script>

<section class="section relative">
  <div class="container-rtl">
    <div class="text-center max-w-2xl mx-auto" use:reveal>
      <span class="eyebrow">مسیر همکاری</span>
      <h2 class="mt-3 text-3xl md:text-4xl font-bold tracking-tight">فرآیند توسعه</h2>
      <p class="lead mt-2 text-slate-600 dark:text-slate-300">۴ گام شفاف تا موفقیت</p>
    </div>

    <ol class="relative mt-8 grid gap-6 md:grid-cols-4 text-center before:hidden md:before:block before:absolute before:inset-x-0 before:top-1/2 before:h-px before:bg-slate-200 dark:before:bg-slate-800" role="list">
      <li class="group relative z-10 card p-6 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm transition-all hover:shadow-md hover:-translate-y-0.5" use:reveal role="listitem">
        <div class="mx-auto grid place-items-center w-10 h-10 rounded-full bg-gradient-to-br from-primary-500/20 to-primary-500/10 text-primary-700 dark:text-primary-300 text-lg font-bold ring-1 ring-primary-500/20">۱</div>
        <h3 class="mt-3 font-semibold">مشاوره و تحلیل نیاز</h3>
      </li>
      <li class="group relative z-10 card p-6 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm transition-all hover:shadow-md hover:-translate-y-0.5" use:reveal role="listitem">
        <div class="mx-auto grid place-items-center w-10 h-10 rounded-full bg-gradient-to-br from-primary-500/20 to-primary-500/10 text-primary-700 dark:text-primary-300 text-lg font-bold ring-1 ring-primary-500/20">۲</div>
        <h3 class="mt-3 font-semibold">طراحی و برنامه‌ریزی</h3>
      </li>
      <li class="group relative z-10 card p-6 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm transition-all hover:shadow-md hover:-translate-y-0.5" use:reveal role="listitem">
        <div class="mx-auto grid place-items-center w-10 h-10 rounded-full bg-gradient-to-br from-primary-500/20 to-primary-500/10 text-primary-700 dark:text-primary-300 text-lg font-bold ring-1 ring-primary-500/20">۳</div>
        <h3 class="mt-3 font-semibold">توسعه و پیاده‌سازی</h3>
      </li>
      <li class="group relative z-10 card p-6 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm transition-all hover:shadow-md hover:-translate-y-0.5" use:reveal role="listitem">
        <div class="mx-auto grid place-items-center w-10 h-10 rounded-full bg-gradient-to-br from-primary-500/20 to-primary-500/10 text-primary-700 dark:text-primary-300 text-lg font-bold ring-1 ring-primary-500/20">۴</div>
        <h3 class="mt-3 font-semibold">تست، تحویل و پشتیبانی</h3>
      </li>
    </ol>
  </div>
</section>
