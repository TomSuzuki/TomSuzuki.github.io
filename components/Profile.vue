<template>
  <div id="profile">
    <div id="profile_title">{{ info.title }}</div>

    <div id="profile_info_outer">
      <div id="profile_info">
        <img id="profile_img_icon" :src="info.profile.img" />
        <div id="profile_table_outer">
          <table id="profile_table">
            <tr v-for="(text, head) in info.profile.table">
              <th>{{ head }}</th>
              <td>{{ text }}</td>
            </tr>
          </table>
        </div>
      </div>
    </div>

    <ul id="profile_sns">
      <li v-for="sns in info.profile.sns">
        <a
          target="_blank"
          class="btn5"
          :class="sns.tag"
          :href="sns.link"
          ><span class="inner_btn5">{{ sns.text }}</span></a
        >
      </li>
    </ul>

    <p id="profile_msg" class="center">{{ info.profile.msg }}</p>
  </div>
</template>

<script>
export default {
  name: "Head",
  props: ["info"],
};
</script>

<style lang="scss">
#profile {
  padding: 0 2rem;
}

#profile_img_header {
  width: 100%;
  height: 640px;
  max-width: 100%;
  object-fit: cover;
  object-position: 50% 50%;
  overflow: hidden;
}

#profile_title {
  font-size: xx-large;
  padding: 0.25em 1em;
  text-align: center;
  margin: 2em 0 0.75em;
}

#profile_info_outer {
  text-align: center;
}

#profile_info {
  display: inline-flex;
  text-align: center;
  justify-content: center;
  flex-flow: wrap;
  align-items: center;
  gap: 1em 3em;
}

#profile_img_icon {
  border-radius: 50%;
  width: 8rem;
  text-align: center;
}

#profile_table_outer {
  text-align: center;
}

#profile_table {
  display: inline-block;
  text-align: center;
  font-size: medium;
  border-spacing: 0.25em;
  margin-top: 0.5em;
}

#profile_table th,
#profile_table td {
  border-radius: var(--default_radius);
  padding: 0.25em 0.5em;
}

#profile_table th {
  text-align: center;
  background-color: var(--color_theme);
  color: var(--color_white);
}

#profile_table td {
  text-align: left;
}

#profile_sns {
  padding: 1em 0;
  list-style-type: none;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

#profile_msg {
  padding: 2em 0 1em;
}
</style>
