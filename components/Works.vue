<template>
  <div id="works">
    <h3>{{ info.works.title }}</h3>
    <div
      class="works_page"
      v-for="(page, index) in pages"
      v-if="index == now_page"
    >
      <div class="works_frame">
        <div v-for="content of page">
          <card :content="content" />
        </div>
      </div>
    </div>
    <div class="works_page_controller">
      <a
        class="works_page_controller_before"
        v-if="this.now_page > 0"
        @click="before_page"
        >前のページへ</a
      >
      <span class="works_page_controller_count"
        >{{ 1 + now_page }} / {{ page_count }}</span
      >
      <a
        class="works_page_controller_next"
        v-if="this.now_page < this.page_count - 1"
        @click="next_page"
        >次のページへ</a
      >
    </div>
  </div>
</template>

<script>
import Card from "~/components/Card.vue";

import works from "@/assets/json/works.json";

// make page
const page_cards = 12;
let pages = [];
let page_count = -1;
for (const i in works["contents"]) {
  // page add
  if (i % page_cards == 0) {
    page_count++;
    pages[page_count] = [];
  }

  // card add
  pages[page_count].push(works["contents"][i]);
}

export default {
  name: "Works",
  props: ["info"],
  data() {
    return {
      pages: pages,
      page_count: page_count + 1,
      now_page: 0,
    };
  },
  methods: {
    next_page: function () {
      this.$scroll_to_id("works", -16);
      this.now_page++;
    },
    before_page: function () {
      this.$scroll_to_id("works", -16);
      this.now_page--;
    },
  },
  components: {
    Card,
  },
};
</script>

<style lang="scss">
.works_frame {
  position: relative;
  padding: 0 0 1.5em;
  width: 100%;
  box-sizing: border-box;
  display: grid;
  grid-auto-flow: row;
  grid-template-columns: repeat(auto-fill, minmax(312px, 1fr));
  grid-auto-rows: auto;
  grid-gap: 20px;
}

.works_page_controller {
  display: grid;
  grid-template-areas: "works_page_controller_before works_page_controller_count works_page_controller_next";
  grid-template-columns: 1fr 1fr 1fr;
  padding: 0 2em;
  box-sizing: border-box;
}

.works_page_controller_count {
  grid-area: works_page_controller_count;
  text-align: center;
}

.works_page_controller_before {
  grid-area: works_page_controller_before;
  text-align: left;
}

.works_page_controller_next {
  grid-area: works_page_controller_next;
  text-align: right;
}
</style>
