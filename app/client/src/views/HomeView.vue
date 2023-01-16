<template>
  <div>
    <div>
      <!--      <span class="text-red-400 text-sm"-->
      <!--        > SOR:burada view.vue dosyasında component'leri böyle sıra ile diziyorum ama-->
      <!--        bu doğru mu?</span-->
      <!--      >-->
      <!--      <span class="text-red-400 text-sm"-->
      <!--        >component isimlendirmesi doğru mu?</span-->
      <!--      >-->
      <TheIntroduction />
      <RecentPosts />
      <FeaturedWorks />
    </div>
  </div>
</template>

<script>
import TheIntroduction from "@/components/Introduction.vue";
import RecentPosts from "@/components/RecentPosts.vue";
import FeaturedWorks from "@/components/FeaturedWorks.vue";

export default {
  name: "HomeView",
  components: { FeaturedWorks, RecentPosts, TheIntroduction },
  setup() {
    return {};
  },
};
</script>

<style scoped></style>
