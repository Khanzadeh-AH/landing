<script lang="ts">
  import { onMount } from 'svelte';
  let hidden = false;
  const KEY = 'tb_sticky_cta_hide';

  onMount(() => {
    try {
      hidden = localStorage.getItem(KEY) === '1';
    } catch (_) {
      // ignore
    }
  });

  function dismiss() {
    hidden = true;
    try {
      localStorage.setItem(KEY, '1');
    } catch (_) {
      // ignore
    }
  }
</script>

{#if !hidden}
  <div class="fixed inset-x-0 bottom-0 z-40 md:hidden">
    <div class="mx-3 mb-3 rounded-2xl border border-slate-200 dark:border-slate-800 bg-white/80 dark:bg-slate-950/70 backdrop-blur shadow-lg">
      <div class="flex items-center justify-between gap-3 p-3">
        <div class="text-sm text-slate-800 dark:text-slate-200">
          آماده شروع هستید؟
          <span class="text-slate-500 dark:text-slate-400">مشاوره رایگان در کمتر از ۲ ساعت پاسخ داده می‌شود.</span>
        </div>
        <div class="flex items-center gap-2">
          <a href="#contact" class="btn-primary px-3 py-2 text-xs">درخواست مشاوره</a>
          <button type="button" aria-label="بستن" class="inline-flex h-8 w-8 items-center justify-center rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800" on:click={dismiss}>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-4 w-4 text-slate-500"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/></svg>
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}
