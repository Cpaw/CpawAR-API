<template>
  <article>
    <section class="challenge">
      <h2>問題追加</h2>
      <form v-on:submit.prevent="addChallenge">
        <div class="field">
          <div class="label">
            <label for="title">Title:</label>
          </div>
          <div class="input">
            <input class="title" type="text" name="title" v-model="challenge.title">
          </div>
        </div>
        <div class="field">
          <div class="label">
            <label for="weight">Weight</label>
          </div>
          <div class="weight">
            <textarea type="text" name="weight" v-model="challenge.weight" />
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
export default {
  data () {
    return {
      challenge: {
        title: '',
        contents: '',
        weight: ''
      }
    }
  }
}
</script>

<style scoped>
</style>
