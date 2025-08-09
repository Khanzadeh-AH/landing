<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
</script>

<section id="home" class="section relative overflow-hidden">
  <div class="pointer-events-none absolute inset-0 -z-10">
    <div class="absolute -top-24 right-[-6rem] h-64 w-64 rounded-full bg-primary-500/20 blur-3xl"></div>
    <div class="absolute -bottom-24 left-[-6rem] h-64 w-64 rounded-full bg-cyan-400/10 blur-3xl"></div>
  </div>

  <div class="container-rtl grid items-center gap-10 md:grid-cols-2">
    <div class="text-right" use:reveal>
      <span class="badge">شرکت توسعه نرم‌افزار و هوش مصنوعی</span>
      <h1 class="mt-4 text-3xl sm:text-4xl md:text-5xl lg:text-6xl leading-tight">
        رشد فروش و بهره‌وری با <span class="bg-gradient-to-l from-primary-500 to-cyan-400 bg-clip-text text-transparent">راهکارهای هوشمند</span> اختصاصی
      </h1>
      <p class="lead mt-4 max-w-prose text-slate-700 dark:text-slate-300">
        در تهران‌بات، تیم ما با ترکیب توسعه نرم‌افزار مدرن و هوش مصنوعی عملیاتی، به شما کمک می‌کند سریع‌تر به نتایج تجاری برسید — با ریسک کمتر و بازگشت سرمایه‌ بالاتر.
      </p>
      <div class="mt-6 flex flex-row-reverse gap-3">
        <a href="#contact" class="btn-primary btn-lg focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500">مشاوره رایگان</a>
        <a href="#services" class="btn-ghost focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500">مشاهده خدمات</a>
      </div>
      <div class="mt-5 flex flex-row-reverse items-center gap-3 text-sm text-slate-600 dark:text-slate-300">
        <div class="inline-flex items-center gap-2 rounded-xl border border-slate-200 dark:border-slate-800 bg-white/60 dark:bg-slate-900/60 px-3 py-2">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-4 w-4 text-emerald-500"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z" clip-rule="evenodd"/></svg>
          <span>مورد اعتماد بیش از ۳۰ کسب‌وکار</span>
        </div>
        <div class="inline-flex items-center gap-2 rounded-xl border border-slate-200 dark:border-slate-800 bg-white/60 dark:bg-slate-900/60 px-3 py-2">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="h-4 w-4 text-primary-500"><path d="M12 2a10 10 0 100 20 10 10 0 000-20zm1 5a1 1 0 10-2 0v6a1 1 0 00.553.894l4 2a1 1 0 10.894-1.788L13 11.382V7z"/></svg>
          <span>میانگین زمان پاسخ‌گویی: کمتر از ۲ ساعت</span>
        </div>
      </div>
    </div>
    <div class="order-first md:order-last" use:reveal>
      <div class="card p-4 sm:p-6">
        <img src="https://images.unsplash.com/photo-1515879218367-8466d910aaa4?q=80&w=1600&auto=format&fit=crop" alt="تصویر برنامه‌نویسی و توسعه نرم‌افزار" width="1600" height="900" decoding="async" fetchpriority="high" class="rounded-xl object-cover w-full h-72 md:h-80 lg:h-96 shadow-md ring-1 ring-slate-200/60 dark:ring-slate-800/60"/>
      </div>
    </div>
  </div>
</section>
