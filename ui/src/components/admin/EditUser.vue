<template>
  <article v-if="isAdmin">
    <section class="challenge">
    </section>
  </article>
</template>

<script>
export default {
  data () {
    return {
      isAdmin: false
    }
  },
  mounted () {
    HTTP.get(`role`,
      {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })
      .then(response => {
        if (response.data.results == 'admin') {
          this.$data.isAdmin = true
        }
      })
      .catch(e => {
      })
  }
}
</script>

<style scoped>
</style>
