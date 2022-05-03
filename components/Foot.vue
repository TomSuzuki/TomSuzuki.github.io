<template>
  <div id="foot">
    <ul id="footer_list">
      <li v-for="msg in info.foot.msg_list" v-html="msg"></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Foot",
  props: ["info"],
};
</script>

<style lang="scss">
#foot {
  background: var(--color_light_gray);
  margin-top: 5em;
  padding: 1em 2em;
  box-sizing: border-box;
}

#footer_list {
  padding: 2em 0;
  margin: 0;
  list-style-type: none;
}

#footer_list li {
  line-height: 2em;
}
</style>
