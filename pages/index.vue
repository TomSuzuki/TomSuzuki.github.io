<template>
  <div>
    <Background/>
    <Head :info="info"/>
    <Profile :info="info"/>
  </div>
</template>

<script>
import '@/assets/css/color.css'
import '@/assets/css/common.css'
import '@/assets/css/button.css'

import text from "@/assets/json/info.json";

export default {
  name: 'IndexPage',
  data(){
    return {
      info: text
    }
  }
}

    // <Profile/>
    // <Works/>
    // <Log/>
    // <Skill/>
    // <Archive/>
    // <Foot/>
</script>
