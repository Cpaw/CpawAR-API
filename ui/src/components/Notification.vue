<template>
  <article>
    <section>
      <ol v-for="notification in notifications" :key="notification.id">
        <li class="title">{{notification.title}}</li>
        <li class="text">{{notification.notification}}</li>
        <li class="date">{{notification.updated_at}}</li>
      </ol>
    </section>
  </article>
</template>
<script>
import {HTTP} from './Header'
export default {
  data () {
    return {
      notifications: {
        id: 0,
        title: '',
        notification: '',
        updated_at: ''
      }
    }
  },
  created () {
    HTTP.get(`notifications`)
      .then(response => {
        this.$data.notifications = response.data.results.notifications
        console.log(this.$data.notifications)
        for (let notification of this.$data.notifications) {
          notification.updated_at = notification.updated_at.slice(0, 19).replace('T', '\n')
        }
        console.log(this.$data.notifications)
      })
      .catch(e => {
      })
  }
}
</script>
<style scoped>
h2 {
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  font-size: 42px;
}
ol {
  background: #FFF;
  border-radius: 5px 5px;
  border: 3px solid #6699cc;
  width: 50vw;
  margin: 3em auto;
  padding-top: 1em;
}
li {
  padding: 7px;
  list-style: none;
}
.title {
  font-size: 36px;
  padding: 0 0 1.2em 0;
}
.text {
  font-size: 22px;
  font-family: "a-otf-ud-shin-maru-go-pr6n";
  padding: 0 0 1.5em 0;
}
.date {
  font-size: 20px;
  text-align: right;
}
</style>
