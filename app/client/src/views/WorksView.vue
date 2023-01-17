<template>
  <div class="container flex justify-center mb-16 md:mb-32">
    <div class="">
      <h1 class="font-extrabold text-[44px] mb-4">Work</h1>
      <FeaturedWorksCard class="mt-10" v-for="i in 3" :key="i" />
    </div>
  </div>
</template>

<script>
import FeaturedWorksCard from "@/components/FeaturedWorksCard.vue";

export default {
  name: "WorksView",
  components: { FeaturedWorksCard },
};
</script>

<style scoped></style>
