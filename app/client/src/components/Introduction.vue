<template>
  <div class="container mb-10">
    <div class="flex flex-col items-center mt-6 md:flex-row-reverse">
      <!--    <span class="text-red-400 text-sm"-->
      <!--      >SOR: burada 2 kısım var, paragraflar ve resim. ben bunlara static w-h-->
      <!--      vermek yerine direkt 1/2 vererek boyurlandırdım. bu doğru mu?</span-->
      <!--    >-->
      <!--    <span class="text-red-400 text-sm"-->
      <!--      >SOR: buradaki img'ye mobile ve desktop için ayrı ayrı mauel w-h verdim.-->
      <!--      bunları direkt static olarak w-h vermek doğru mu?</span-->
      <!--    >-->
      <!--    <span class="text-red-400 text-sm"-->
      <!--      >SOR: tasarım dosyasını abdülkadir abiye yorumla?</span-->
      <!--    >-->
      <div
        class="flex justify-center items-center md:items-start w-full md:w-1/2 md:h-1/2"
      >
        <img
          class="w-40 h-40 md:w-64 md:h-64 object-cover rounded-full shadow shadow-2xl"
          src="../img/me.jpeg"
          alt=""
        />
      </div>
      <div class="text-center md:text-start">
        <h1
          class="font-bold text-3xl mt-6 md:text-[44px] md:leading-[60px] letter-spacing"
        >
          Hi, I am Talha, Candidate of Computer Engineering
        </h1>
        <p class="mt-6 leading-6 text-colortext md:mt-8">
          Türkiye'nin Isparta ilinde bilgisayar mühendisliği okuyorum. 2. sınıf
          öğrencisiyim. Şuanda web üzerine çalışıyorum.
        </p>
        <button type="button" class="button bg-lightred mt-6 md:mt-8">
          Download Resume
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "TheIntroduction",
};
</script>

<style scoped></style>
