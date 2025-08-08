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
      <h1 class="mt-4 text-3xl sm:text-4xl md:text-5xl leading-tight">
        راهکارهای هوشمند برای رشد کسب‌وکار شما
      </h1>
      <p class="lead mt-4 max-w-prose">
        تهران‌بات با تجربه در توسعه نرم‌افزارهای مدرن و سیستم‌های هوش مصنوعی، به شما کمک می‌کند تا سریع‌تر، هوشمندتر و بهینه‌تر حرکت کنید.
      </p>
      <div class="mt-6 flex flex-row-reverse gap-3">
        <a href="#services" class="btn-primary">مشاهده خدمات</a>
        <a href="#contact" class="btn-ghost">شروع پروژه</a>
      </div>
    </div>
    <div class="order-first md:order-last" use:reveal>
      <div class="card p-4 sm:p-6">
        <img src="https://images.unsplash.com/photo-1515879218367-8466d910aaa4?q=80&w=1600&auto=format&fit=crop" alt="تصویر برنامه‌نویسی و توسعه نرم‌افزار" width="1600" height="900" decoding="async" fetchpriority="high" class="rounded-lg object-cover w-full h-64"/>
      </div>
    </div>
  </div>
</section>
