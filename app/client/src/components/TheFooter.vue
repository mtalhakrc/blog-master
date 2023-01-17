<template>
  <div class="container mb-12">
    <div class="flex flex-col items-center">
      <div class="flex flex-row">
        <a href=""><InstagramSVG style="margin-right: 30px" /></a>
        <a href=""><FacebookSVG style="margin-right: 30px" /></a>
        <a href=""><LinkedinSVG style="margin-right: 30px" /></a>
        <a href=""><TwitterSVG /></a>
      </div>
      <div class="text-colortext font-normal text-[14px] mt-6">
        Copyright ©2020 All rights reserved
      </div>
    </div>
  </div>
</template>

<script>
import InstagramSVG from "../svg/instagram.svg";
import FacebookSVG from "../svg/facebook.svg";
import LinkedinSVG from "../svg/linkedin.svg";
import TwitterSVG from "../svg/twitter.svg";

export default {
  name: "TheFooter",
  components: {
    InstagramSVG,
    FacebookSVG,
    LinkedinSVG,
    TwitterSVG,
  },
};
</script>

<style scoped></style>
