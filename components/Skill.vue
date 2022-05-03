<template>
  <div id="skill">
    <h3>{{ info.skill.title }}</h3>
    <div v-for="skill in skills">
      <h4>{{ skill.title }}</h4>
      <ul class="skill_list">
        <li v-for="s in skill.contents" :class="s.tag">{{ s.name }}</li>
      </ul>
    </div>
  </div>
</template>

<script>
import skills from "@/assets/json/skill.json";

export default {
  name: "Log",
  props: ["info"],
  data() {
    return {
      skills: skills["skill"],
    };
  },
};
</script>

<style lang="scss">
.skill {
  display: block;
  position: relative;
  max-width: 100%;
}

.skill_list {
  width: 100%;
  display: inline-flex;
  flex-wrap: wrap;
  height: auto;
  list-style: none;
  box-sizing: border-box;
  padding: 0 0.8em 0.8em;
}

.skill_list > li {
  min-width: 9em;
  height: auto;
  box-sizing: border-box;
  border-radius: var(--default_radius);
  padding: 0.6em 1em;
  text-align: center;
  display: inline-block;
  font-weight: bold;
  background: var(--color_white2);
  text-decoration: none;
  margin: 0.4em 0.8em;
  border-style: solid;
  border-width: 2px;
  border-left-width: 0.55em;
  word-wrap: none;
  white-space: pre;
}
</style>
