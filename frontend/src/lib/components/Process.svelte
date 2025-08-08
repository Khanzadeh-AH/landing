<script lang="ts">
  import { reveal } from '$lib/actions/reveal';
</script>

<section class="section">
  <div class="container-rtl">
    <div class="text-center" use:reveal>
      <h2 class="text-3xl md:text-4xl">فرآیند توسعه</h2>
      <p class="lead mt-2">۴ گام شفاف تا موفقیت</p>
    </div>

    <ol class="mt-8 grid gap-6 md:grid-cols-4 text-center">
      <li class="card p-6" use:reveal>
        <div class="text-2xl font-bold">۱</div>
        <h3 class="mt-2 font-semibold">مشاوره و تحلیل نیاز</h3>
      </li>
      <li class="card p-6" use:reveal>
        <div class="text-2xl font-bold">۲</div>
        <h3 class="mt-2 font-semibold">طراحی و برنامه‌ریزی</h3>
      </li>
      <li class="card p-6" use:reveal>
        <div class="text-2xl font-bold">۳</div>
        <h3 class="mt-2 font-semibold">توسعه و پیاده‌سازی</h3>
      </li>
      <li class="card p-6" use:reveal>
        <div class="text-2xl font-bold">۴</div>
        <h3 class="mt-2 font-semibold">تست، تحویل و پشتیبانی</h3>
      </li>
    </ol>
  </div>
</section>
