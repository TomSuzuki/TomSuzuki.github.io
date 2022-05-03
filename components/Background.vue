<template>
  <div id="background"></div>
</template>

<script>
export default {
  name: "Background",
};
</script>

<style>
#background {
  /* --- background pattern --- */
  position: fixed;
  min-width: 100%;
  min-height: 100%;
  width: calc(100% + 20px);
  height: calc(100% + 20px);
  top: 0;
  left: 0;
  overflow: hidden;
  z-index: -1;
  background: var(--Color_White4);
  background-image: linear-gradient(var(--color_white5) 1px, transparent 0),
    linear-gradient(90deg, var(--color_white5) 1px, transparent 0);
  background-size: 20px 20px;

  /* --- for animation --- */
  animation-duration: 1.2s;
  animation-name: background_animation;
  animation-timing-function: linear;
  animation-iteration-count: infinite;
}

@keyframes background_animation {
  0% {
    transform: translateX(-20px) translateY(-20px);
  }

  100% {
    transform: translateX(0) translateY(0);
  }
}
</style>
