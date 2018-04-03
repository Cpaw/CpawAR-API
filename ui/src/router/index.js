import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Challenges from '@/components/Challenges'
import Challenge from '@/components/Challenge'
import Ranking from '@/components/Ranking'
import Auth from '@/components/Auth'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/challenges',
      name: 'challenges',
      component: Challenges
    },
    {
      path: '/challenge/:challenge_id',
      name: 'challenge',
      component: Challenge
    },
    {
      path: '/ranking',
      name: 'ranking',
      component: Ranking
    },
    {
      path: '/auth',
      name: 'auth',
      component: Auth
    }
  ]
})
