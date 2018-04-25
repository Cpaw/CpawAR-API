<template>
  <article>
    <section class="challenge">
      <router-link :to="{ name: 'addchallenge' }">Add Challenge</router-link>
      <router-link :to="{ name: 'editchallenge' }">Edit Challenge</router-link>
    </section>
     <section class="user">
      <router-link :to="{ name: 'edituser' }"></router-link>
    </section>
  </article>
</template>

<script>
export default {
  data () {
  }
}
</script>

<style scoped>
</style>
