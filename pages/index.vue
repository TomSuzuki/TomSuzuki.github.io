<template>
  <div>
    <Background />
    <Head :info="info" />
    <Profile :info="info" />
    <div id="main-contents">
      <Works :info="info" />
      <Log :info="info" />
      <Skill :info="info" />
      <Archive :info="info" />
    </div>
    <Foot :info="info" />
  </div>
</template>

<script>
import "@/assets/css/color.css";
import "@/assets/css/common.css";
import "@/assets/css/button.css";

import text from "@/assets/json/info.json";

export default {
  name: "IndexPage",
  data() {
    return {
      info: text,
    };
  },
};
</script>

<style>
#main-contents {
  padding: 0 2em;
  box-sizing: border-box;

  width: var(--px_tablet);
  min-width: 80%;
  max-width: 100%;
  margin: 0 auto;
  color: var(--color_text);
}
</style>
