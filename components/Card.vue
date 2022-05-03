<template>
  <div class="works_content_frame">
    <img
      :src="'./img/' + content.img"
      alt=""
      onerror="this.src='./img/default.gif';"
    />
    <div class="works_info_frame">
      <div class="works_title">{{ content.title }}</div>
      <span class="works_date">{{ content.date }}</span>
      <span class="works_link" v-if="content.link">
        ｜<a :href="content.link.link">{{ content.link.text }}</a>
      </span>
      <div class="works_text" v-html="content.text"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Card",
  props: ["content"],
};
</script>

<style lang="scss">
.works_content_frame {
  box-sizing: border-box;
  display: grid;
  grid-template-rows: 8em auto;
  grid-template-columns: 100%;
  position: relative;
  background: var(--color_white);
  border-radius: var(--default_radius_large);
}

.works_content_frame img {
  box-sizing: border-box;
  border-top-left-radius: var(--default_radius_large);
  border-top-right-radius: var(--default_radius_large);
  border: 2px solid var(--color_gray);
  border-bottom: none;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: 50% 50%;
  overflow: hidden;
}

.works_info_frame {
  box-sizing: border-box;
  width: 100%;
  padding: 0 0.5em;
  position: relative;
  border: 2px solid var(--color_gray);
  border-top: none;
  border-bottom-left-radius: var(--default_radius_large);
  border-bottom-right-radius: var(--default_radius_large);
  text-align: center;
}

.works_title {
  font-size: large;
  margin: 0.5em 0 0;
}

.works_date {
  padding-left: 1.25em;
  font-size: small;
  color: var(--color_dark_gray);
}

.works_link {
  box-sizing: border-box;
  max-width: 100%;
  font-size: small;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

.works_text {
  font-size: small;
  padding: 1em 1em;
  text-align: justify;
}
</style>
