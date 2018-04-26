<template>
  <article v-if="isAdmin">
    <h2>問題追加</h2>
    <section v-for="challenge in challenges" :key="challenge.id">
      <form v-on:submit.prevent="editChallenge">
        <div class="field">
          <div class="label">
            <label for="title">Title:</label>
          </div>
          <div class="input">
            <input class="title" type="text" name="title" v-model="challenge.Title">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="contents">Challenge text:</label>
          </div>
          <div class="contents">
            <textarea type="text" name="contents" v-model="challenge.QuestionText" cols="60" rows="5" />
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="weight">Weight:</label>
          </div>
          <div class="weight">
            <input type="number" name="weight" v-model="challenge.Weight" />
          </div>
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
      <div class="error" v-if="addChallengeError">Invalid data</div>
    </section>
  </article>
</template>

<script>
import {HTTP} from '../Header'

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
        if (response.data.results === 'admin') {
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
