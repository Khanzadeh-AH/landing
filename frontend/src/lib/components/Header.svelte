<script lang="ts">
  let open = false;
  const nav = [
    { href: '#home', label: 'خانه' },
    { href: '#services', label: 'خدمات' },
    { href: '#portfolio', label: 'نمونه کارها' },
    { href: '#about', label: 'درباره ما' },
    { href: '#contact', label: 'تماس با ما' }
  ];
</script>

<header class="sticky top-0 z-40 border-b border-slate-200/70 dark:border-slate-800/70 bg-white/70 dark:bg-slate-950/60 backdrop-blur">
  <div class="container-rtl flex h-16 items-center justify-between">
    <a href="#home" class="flex items-center gap-2 text-lg font-extrabold">
      <span class="inline-flex h-8 w-8 items-center justify-center rounded-lg bg-primary-600 text-white">TB</span>
      <span>تهران‌بات</span>
    </a>

    <nav class="hidden md:flex items-center gap-6 text-sm">
      {#each nav as item}
        <a href={item.href} class="text-slate-700 dark:text-slate-200 hover:text-primary-600">{item.label}</a>
      {/each}
      <a href="#contact" class="btn-primary">مشاوره رایگان</a>
    </nav>

    <button class="md:hidden inline-flex items-center justify-center rounded-xl p-2 hover:bg-slate-100 dark:hover:bg-slate-800" on:click={() => (open = !open)} aria-label="toggle menu">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
      </svg>
    </button>
  </div>

  {#if open}
    <div class="md:hidden border-t border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-950">
      <div class="container-rtl py-4 flex flex-col gap-3">
        {#each nav as item}
          <a href={item.href} class="py-2" on:click={() => (open = false)}>{item.label}</a>
        {/each}
        <a href="#contact" class="btn-primary" on:click={() => (open = false)}>مشاوره رایگان</a>
      </div>
    </div>
  {/if}
</header>
