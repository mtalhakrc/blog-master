<template>
  <!--  <span class="text-red-400 text-sm"-->
  <!--    > mesela burada header biraz daha aşağıda olsun diye mt verdim, bu doğru-->
  <!--    mu?</span-->
  <!--  >-->
  <header class="h-12">
    <div class="container h-full flex justify-end items-center">
      <button class="block md:hidden" @click="isMenushow = !isMenushow">
        <menuSVG />
      </button>

      <div
        class="flex flex-col bg-white border text-lg font-semibold tracking-tight z-10 absolute left-0 top-full w-full p-4 opacity-0 transition pointer-events-none group md:flex-row md:static md:opacity-100 md:pointer-events-auto md:justify-end md:border-0 md:p-0 md:top-0 md:mt-10"
        :class="{
          'opacity-100': isMenushow,
          'pointer-events-auto': isMenushow,
        }"
      >
        <RouterLink to="/works" class="menu-item">Works</RouterLink>
        <RouterLink to="/menu" class="menu-item">Menu</RouterLink>
        <RouterLink to="/contact" class="menu-item">Contact</RouterLink>
      </div>
    </div>
  </header>
</template>

<script>
import menuSVG from "@/svg/menu.svg";
import { ref } from "vue";

export default {
  name: "TheHeader",
  components: {
    menuSVG,
  },
  setup() {
    const isMenushow = ref(false);
    return {
      isMenushow,
    };
  },
};
</script>

<style scoped>
.menu-item {
  @apply inline-flex h-10 items-center justify-center items-center w-full mt-2  transition  md:w-auto md:px-4 md:hover:bg-gray-400 md:mt-0;
}
</style>
