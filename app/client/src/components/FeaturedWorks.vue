<template>
  <div class="container mt-6 mb-16 md:mb-32">
    <div class="flex flex-col items-center md:items-start">
      <h2 class="text-lg">Featured Works</h2>
      <FeaturedWorksCard class="mt-4" />
    </div>
  </div>
</template>

<script>
import FeaturedWorksCard from "@/components/FeaturedWorksCard.vue";

export default {
  name: "FeaturedWorks",
  components: { FeaturedWorksCard },
};
</script>

<style scoped></style>
