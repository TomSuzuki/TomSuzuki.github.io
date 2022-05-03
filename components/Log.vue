<template>
  <div id="Log">
    <h3>{{ info.log.title }}</h3>
    <ul class="log_list">
      <li v-for="log in logs">
        <div class="log_date">{{ log.date }}</div>
        <div class="log_content">
          <div class="log_title" v-html="log.title"></div>
          <div v-if="log.text" class="log_text" v-html="log.text"></div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import logs from "@/assets/json/log.json";

export default {
  name: "Log",
  props: ["info"],
  data() {
    return {
      logs: logs["log"],
    };
  },
};
</script>

<style lang="scss">
.log_list {
  list-style: none;
}

.log_title {
  font-size: medium;
  font-weight: bold;
  padding: 0.5em 0;
}

.log_text {
  font-size: small;
  padding-bottom: 1.5em;
  white-space: pre-wrap;
}

@media (min-width: 600px) {
  .log_list li {
    overflow: hidden;
    margin: 0;
    position: relative;
    display: flex;
  }

  .log_date {
    width: 80px;
    float: left;
    padding-top: 12px;
  }

  .log_text {
    padding-left: 1em;
  }

  .log_content {
    width: calc(100% - 80px);
    left: 0;
    float: left;
    border-left: 3px var(--color_gray) solid;
    padding-left: 30px;
    box-sizing: border-box;
  }

  .log_content:before {
    content: "";
    width: 11.5px;
    height: 11.5px;
    background: var(--color_theme);
    position: absolute;
    left: 76px;
    top: 16px;
    border-radius: 100%;
  }
}

@media (max-width: 599px) {
  .log_list li {
    margin: 0;
    position: relative;
    padding-left: 14px;
    border-left: 3px var(--color_gray) solid;
  }

  .log_date {
    padding-top: 12px;
  }

  .log_title {
    padding-top: 2px;
    padding-left: 2px;
  }

  .log_text {
    padding-left: 1em;
  }

  .log_content:before {
    content: "";
    width: 11px;
    height: 11px;
    background: var(--color_theme);
    position: absolute;
    left: -7px;
    top: 18px;
    border-radius: 100%;
  }
}
</style>
