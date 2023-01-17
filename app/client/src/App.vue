<template>
  <div>
    <TheHeader />
  </div>
  <RouterView />
  <div>
    <TheFooter/>
  </div>


  <!--  <div class="container">-->
  <!--  </div>-->
</template>
<script>
import { defineComponent } from "vue";
import TheHeader from "@/components/Header.vue";
import TheFooter from "@/components/TheFooter.vue";

export default defineComponent({
  components: { TheFooter, TheHeader },
});
</script>

<style scoped></style>
