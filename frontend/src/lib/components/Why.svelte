<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
  const items = [
    'تجربه در پروژه‌های متنوع',
    'استفاده از جدیدترین فناوری‌ها',
    'پشتیبانی ۲۴/۷',
    'تحویل به‌موقع پروژه‌ها',
    'قیمت‌گذاری منصفانه و شفاف',
    'تخصص فنی بالا'
  ];
</script>

<section id="about" class="section relative overflow-hidden">
  <!-- Decorative gradients -->
  <div class="pointer-events-none absolute -z-10 -top-10 -left-10 h-44 w-44 rounded-full bg-primary-500/10 blur-2xl dark:opacity-70" aria-hidden="true"></div>
  <div class="pointer-events-none absolute -z-10 -bottom-12 -right-8 h-56 w-56 rounded-full bg-cyan-400/10 blur-3xl dark:opacity-70" aria-hidden="true"></div>
  <div class="container-rtl">
    <div class="text-center max-w-2xl mx-auto" use:reveal>
      <span class="eyebrow">چرا ما</span>
      <h2 class="mt-3 text-3xl md:text-4xl font-bold tracking-tight">چرا تهران‌بات؟</h2>
      <p class="lead mt-2 text-slate-600 dark:text-slate-300">ترکیبی از تجربه، تخصص و تعهد</p>
    </div>

    <div class="mt-8 grid gap-6 sm:grid-cols-2 lg:grid-cols-3" role="list">
      {#each items as item}
        <article class="group card p-6 flex items-center gap-3 rounded-2xl border border-slate-200/60 dark:border-slate-700/60 bg-white/60 dark:bg-slate-900/40 backdrop-blur shadow-sm transition-all ring-1 ring-transparent hover:ring-primary-500/10" use:reveal role="listitem">
          <div class="w-9 h-9 rounded-lg bg-gradient-to-br from-primary-500/15 to-primary-500/5 text-primary-600 dark:text-primary-400 grid place-items-center transition-transform duration-200 group-hover:scale-110">
            <svg class="w-5 h-5" viewBox="0 0 20 20" fill="currentColor" xmlns="http://www.w3.org/2000/svg" aria-hidden="true"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/></svg>
          </div>
          <span class="text-slate-700 dark:text-slate-200">{item}</span>
        </article>
      {/each}
    </div>
  </div>
</section>
