<template>
  <div class="archive">
    <h3>{{ info.archive.title }}</h3>
    <details v-for="archive in archives" class="archive_details">
      <summary>{{ archive.title }} ({{ archive.contents.length }})</summary>
      <ul>
        <li v-for="content in archive.contents">
          <a :href="content.link">{{ content.name }}</a>
        </li>
      </ul>
    </details>
  </div>
</template>

<script>
import archives from "@/assets/json/archive.json";

export default {
  name: "Archive",
  props: ["info"],
  data() {
    return {
      archives: archives["archive"],
    };
  },
};
</script>

<style lang="scss">
.archive_details {
  padding: 0 1em;
  box-sizing: border-box;
}

.archive_details ul {
  margin-top: 0.25em;
}

.archive_details summary {
  font-size: small;
  width: auto;
  display: inline;
  user-select: none;
}

.archive_details summary:focus {
  outline: none;
}

.archive_details summary:hover {
  color: var(--color_theme);
  cursor: pointer;
}
</style>
