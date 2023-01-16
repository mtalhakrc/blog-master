<template>
  <div
    class="flex flex-col max-w-md md:max-w-none items-center md:items-start md:flex-row max-w-lg md:max-w-none border-b-2 relative pb-4 mb-40 w-full"
  >
    <img
      class="w-full md:mr-4 rounded-md md:w-80 float-right"
      style=""
      src="../img/foo.jpg"
      alt="works"
    />
    <div class="flex flex-col w-full">
      <h3 class="font-bold text-2xl mt-2 md:mt-0">Designing Dashboards</h3>
      <div class="mt-5">
        <span class="year mr-4">2020</span>
        <span class="tag font-medium text-lightgray">Dashboard</span>
      </div>
      <p class="mt-4 font-medium color-textcolor leading-6">
        ÇOK FAZLA YAZI OLUNCA SARMALASIN Amet Lorem ipsum dolor sit amet,
        consectetur adipisicing elit. Ea et eveniet fugit id laboriosam
        molestias neque nostrum quaerat quidem sint. Amet Lorem ipsum dolor sit
        amet, consectetur adipisicing elit. Ea et eveniet fugit id laboriosam
        molestias neque nostrum quaerat quidem sint. Amet Lorem ipsum dolor sit
        amet, consectetur adipisicing elit. Ea et eveniet fugit id laboriosam
        molestias neque nostrum quaerat quidem sint. Amet Lorem ipsum dolor sit
        amet, consectetur adipisicing elit. Ea et eveniet fugit id laboriosam
        molestias neque nostrum quaerat quidem sint. Amet Lorem ipsum dolor sit
        amet, consectetur adipisicing elit. Ea et eveniet fugit id laboriosam
        molestias neque nostrum quaerat quidem sint. Amet Lorem ipsum dolor sit
        amet, consectetur adipisicing elit. Ea et eveniet fugit id laboriosam
        molestias neque nostrum quaerat quidem sint.
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: "FeaturedWorksCard",
};
</script>

<style scoped>
.year {
  @apply inline-block bg-black text-white pl-3 pr-3 font-extrabold w-auto  rounded-full;
}
</style>
