<template>
  <div>
    <TheHeader />
  </div>
  <!--  <div class="container">-->
  <RouterView />
  <!--  </div>-->
</template>
<script>
import { defineComponent } from "vue";
import TheHeader from "@/components/Header.vue";

export default defineComponent({
  components: { TheHeader },
});
</script>

<style scoped></style>
