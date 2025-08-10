<script lang="ts">
  import Services from '$lib/components/Services.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import StickyCTA from '$lib/components/StickyCTA.svelte';
  import { reveal } from '$lib/actions/reveal';
</script>

<svelte:head>
  <title>خدمات | تهران‌بات</title>
  <meta name="description" content="لیست خدمات تهران‌بات: توسعه نرم‌افزار وب و موبایل، ربات‌های هوش مصنوعی، RAG، DevOps و اتوماسیون." />
</svelte:head>

<main id="main-content" class="section">
  <div class="container-rtl" use:reveal>
    <h1 class="text-3xl md:text-4xl">خدمات تهران‌بات</h1>
    <p class="lead mt-2">در این صفحه می‌توانید جزئیات خدمات ما را مشاهده کنید.</p>
  </div>

  <section class="section pt-6" aria-labelledby="services-heading">
    <div class="container-rtl">
      <h2 id="services-heading" class="sr-only">خدمات</h2>
      <Services />
    </div>
  </section>

  <section class="section" aria-labelledby="cta-heading">
    <div class="container-rtl" use:reveal>
      <h2 id="cta-heading" class="text-2xl md:text-3xl">برای شروع آماده‌اید؟</h2>
      <p class="mt-2">برای مشاوره رایگان به صفحه تماس مراجعه کنید.</p>
      <div class="mt-4">
        <a href="/contact" class="btn-primary">درخواست مشاوره</a>
      </div>
    </div>
  </section>
</main>

<Footer />
<StickyCTA />
