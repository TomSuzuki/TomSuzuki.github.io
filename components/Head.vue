<!-- Head: (logo, menu) -->

<template>
  <div id="head">
    <div class="logo">
      <span class="logo_text">
        <a :href="info.link">{{ info.title }}</a>
      </span>
    </div>
    <ul id="menu">
      <li v-for="(id, text) in info.menu">
        <a :href="id">{{ text }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Head",
  props: ["info"]
};
</script>

<style lang="scss">
:root {
  --menu_height: 3rem;
}

@media (max-width: 599px) {
  #menu {
    display: none;
  }
}

#head {
  display: grid;
  grid-template-columns: 1fr auto;
  width: 100%;
  position: relative;
  height: var(--menu_height);
  background-color: var(--color_black);

  .logo {
    color: var(--color_white);

    .logo_text {
      display: inline-block;
      height: 100%;
      background-color: var(--color_black2);

      a {
        white-space: nowrap;
        text-decoration: none;
        font-size: large;
        color: var(--color_white);
        padding: 0.2em 1em 0;
        line-height: var(--menu_height);
      }
    }
  }

  #menu {
    text-align: right;
    list-style: none;
    font-size: 0;
    overflow: hidden;
    padding: 0;

    li {
      width: 140px;
      text-align: center;
      display: inline-block;
      font-size: 16px;

      a {
        display: inline-block;
        width: 100%;
        height: 100%;
        background-color: var(--color_black);
        text-decoration: none;
        color: var(--color_white);
        font-weight: bold;
        line-height: 3em;
      }

      a:hover {
        cursor: pointer;
        background-color: var(--color_black2);
      }
    }
  }
}
</style>
