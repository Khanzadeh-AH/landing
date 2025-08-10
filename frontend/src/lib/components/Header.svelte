<script lang="ts">
  let open = false;
  let scrolled = false;
  const nav = [
    { href: '/#home', label: 'خانه' },
    { href: '/#services', label: 'خدمات' },
    { href: '/#portfolio', label: 'نمونه کارها' },
    { href: '/#about', label: 'درباره ما' },
    { href: '/blog', label: 'بلاگ' },
    { href: '/#contact', label: 'تماس با ما' }
  ];
</script>

<svelte:window on:scroll={() => (scrolled = scrollY > 4)} />

<header class={`sticky top-0 z-40 md:z-50 border-b border-slate-200/70 dark:border-slate-800/70 backdrop-blur transition-colors ${scrolled ? 'bg-white/90 dark:bg-slate-950/70 shadow-sm' : 'bg-white/60 dark:bg-slate-950/40'}`}>
  <div class="container-rtl flex h-16 items-center justify-between">
    <a href="/#home" class="flex items-center gap-2 text-lg font-extrabold focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 rounded-lg">
      <span class="inline-flex h-8 w-8 items-center justify-center rounded-lg bg-primary-600 text-white">TB</span>
      <span>تهران‌بات</span>
    </a>

    <nav id="primary-navigation" class="hidden md:flex items-center gap-6 text-sm" aria-label="منوی اصلی">
      {#each nav as item}
        <a href={item.href} class="group relative text-slate-700 dark:text-slate-200 hover:text-primary-600 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 rounded-md">
          <span>{item.label}</span>
          <span aria-hidden="true" class="absolute -bottom-1 right-0 left-0 h-px scale-x-0 origin-center bg-primary-600 transition-transform group-hover:scale-x-100"></span>
        </a>
      {/each}
      <a href="/#contact" class="btn-primary">مشاوره رایگان</a>
      <span class="hidden lg:inline-flex items-center gap-1 rounded-full border border-amber-300/60 bg-amber-50/70 text-amber-800 dark:border-amber-300/20 dark:bg-amber-400/10 dark:text-amber-200 px-2 py-1 text-[11px]">
        ظرفیت این ماه محدود است
      </span>
    </nav>

    <button type="button" class="md:hidden inline-flex items-center justify-center rounded-xl p-2 hover:bg-slate-100 dark:hover:bg-slate-800 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500" on:click={() => (open = !open)} aria-label="باز و بسته کردن منو" aria-expanded={open} aria-controls="mobile-navigation">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
      </svg>
    </button>
  </div>

  {#if open}
    <div id="mobile-navigation" class="md:hidden border-t border-slate-200 dark:border-slate-800 bg-white/95 dark:bg-slate-950/95 backdrop-blur">
      <div class="container-rtl py-4 flex flex-col gap-1">
        {#each nav as item}
          <a href={item.href} class="py-3 rounded-lg hover:bg-slate-100/70 dark:hover:bg-slate-800/60 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500" on:click={() => (open = false)}>{item.label}</a>
        {/each}
        <a href="/#contact" class="btn-primary" on:click={() => (open = false)}>مشاوره رایگان</a>
        <div class="text-xs text-amber-600 dark:text-amber-300 mt-2">میانگین پاسخ‌گویی: کمتر از ۲ ساعت</div>
      </div>
    </div>
  {/if}
</header>
