<script lang="ts">
  import Header from '$lib/components/Header.svelte';
  import Hero from '$lib/components/Hero.svelte';
  import Services from '$lib/components/Services.svelte';
  import Why from '$lib/components/Why.svelte';
  import Portfolio from '$lib/components/Portfolio.svelte';
  import Process from '$lib/components/Process.svelte';
  import Testimonials from '$lib/components/Testimonials.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import StickyCTA from '$lib/components/StickyCTA.svelte';
  import { enhance } from '$app/forms';
  import { page } from '$app/stores';
  import { reveal } from '$lib/actions/reveal';

  const form = $derived($page.form as any);
  const values = $derived(form?.values || {});
</script>

<Header />
<main id="main-content">
  <Hero />
  <Services />
  <Why />
  <Portfolio />
  <Process />
  <Testimonials />

  <section id="contact" class="section">
    <div class="container-rtl grid gap-10 md:grid-cols-2 items-start">
      <div use:reveal>
        <h2 class="text-3xl md:text-4xl">ارتباط با ما</h2>
        <p class="lead mt-2">برای مشاوره رایگان و شروع همکاری، فرم زیر را تکمیل کنید یا از روش‌های زیر استفاده کنید.</p>
        <ul class="mt-4 grid gap-3 text-sm text-slate-700 dark:text-slate-300">
          <li class="inline-flex items-center gap-2"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-4 w-4 text-emerald-500"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd"/></svg>مشاوره واقعی و بدون تعهد</li>
          <li class="inline-flex items-center gap-2"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-4 w-4 text-emerald-500"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd"/></svg>میانگین پاسخ‌گویی کمتر از ۲ ساعت</li>
          <li class="inline-flex items-center gap-2"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-4 w-4 text-emerald-500"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd"/></svg>حفظ محرمانگی اطلاعات شما</li>
        </ul>
        <ul class="mt-6 space-y-3 text-slate-700 dark:text-slate-300">
          <li>شماره تماس: <a class="hover:underline ltr" href="tel:+989190737241">+98 919 073 7241</a></li>
          <li>واتساپ: <a class="hover:underline ltr" href="https://wa.me/989190737241" target="_blank" rel="noopener">+98 919 073 7241</a></li>
          <li>تلگرام: <a class="hover:underline" href="https://t.me/Khanzadeh_AH" target="_blank" rel="noopener">Khanzadeh_AH</a></li>
          <li>ایمیل: <a class="hover:underline" href="mailto:khanzadeh78ah@gmail.com">khanzadeh78ah@gmail.com</a></li>
        </ul>
      </div>

      <form method="post" class="card p-6 space-y-4" use:enhance use:reveal>
        {#if form?.success}
          <div role="status" aria-live="polite" class="rounded-xl bg-green-50 text-green-800 dark:bg-green-900/30 dark:text-green-200 p-3 text-sm">پیام شما با موفقیت ارسال شد. به زودی با شما تماس می‌گیریم.</div>
        {/if}
        <div>
          <label class="block text-sm mb-1" for="name">نام و نام خانوادگی</label>
          <input id="name" name="name" value={values.name || ''} autocomplete="name" aria-invalid={form?.errors?.name ? 'true' : 'false'} aria-describedby="name-error" class="w-full rounded-xl border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900" required />
          {#if form?.errors?.name}
            <div id="name-error" class="mt-1 text-xs text-red-600">{form.errors.name}</div>
          {/if}
        </div>
        <div>
          <label class="block text-sm mb-1" for="email">ایمیل</label>
          <input id="email" name="email" type="email" value={values.email || ''} autocomplete="email" aria-invalid={form?.errors?.email ? 'true' : 'false'} aria-describedby="email-error" class="w-full rounded-xl border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900" required />
          {#if form?.errors?.email}
            <div id="email-error" class="mt-1 text-xs text-red-600">{form.errors.email}</div>
          {/if}
        </div>
        <div>
          <label class="block text-sm mb-1" for="phone">شماره تماس</label>
          <input id="phone" name="phone" inputmode="tel" autocomplete="tel" value={values.phone || ''} aria-invalid={form?.errors?.phone ? 'true' : 'false'} aria-describedby="phone-error" class="w-full rounded-xl border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900" />
          {#if form?.errors?.phone}
            <div id="phone-error" class="mt-1 text-xs text-red-600">{form.errors.phone}</div>
          {/if}
        </div>
        <div>
          <label class="block text-sm mb-1" for="message">پیام شما</label>
          <textarea id="message" name="message" rows="5" aria-describedby="message-error" class="w-full rounded-xl border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900">{values.message || ''}</textarea>
          {#if form?.errors?.message}
            <div id="message-error" class="mt-1 text-xs text-red-600">{form.errors.message}</div>
          {/if}
        </div>
        <div class="flex flex-col gap-2 items-end">
          <button class="btn-primary" type="submit">ارسال پیام</button>
          <div class="text-[11px] text-slate-500 dark:text-slate-400">با ارسال فرم، هیچ تعهدی ایجاد نمی‌شود.</div>
        </div>
      </form>
    </div>
  </section>
</main>

<Footer />
<StickyCTA />
