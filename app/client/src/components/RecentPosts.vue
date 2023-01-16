<template>
  <!--  <span class="text-red-400 text-sm"-->
  <!--    >burada container'ı override edip w100% şeklinde bg vermek istiyorum. bunu-->
  <!--    nasıl yapmalıyım?</span-->
  <!--  >-->
  <div class="bg-lightblue py-5 bg-lightblue">
    <div class="container">
      <div class="flex items-center justify-center md:justify-between">
        <div class="child text-lg md:text-xl text-center">Recent posts</div>
        <RouterLink
          class="hidden text-shadow md:inline-block text-darkblue underline hover:no-underline"
          to="/"
          >View all</RouterLink
        >
      </div>

      <!--      <span class="text-red-400 text-sm"-->
      <!--        >SOR:mesela burada da tasarıma göre her bir kutu %45 falan olması lazım-->
      <!--        ama buraya öyle %45 falan verince saçma sapan bir görüntü ortaya-->
      <!--        çıkıyor. bu kutuların boyutlandırması nasıl olmalıo?</span-->
      <!--      >-->
      <div class="flex flex-col md:flex-row">
        <RecentPostsCard class="item" v-for="i in 2" :key="i" />
      </div>
    </div>
  </div>
</template>

<script>
import RecentPostsCard from "@/components/RecentPostsCard.vue";

export default {
  name: "RecentPosts",
  components: { RecentPostsCard },
};
</script>

<style scoped>
.item {
  @apply mt-4;
}
.text-shadow {
  text-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
}
</style>
