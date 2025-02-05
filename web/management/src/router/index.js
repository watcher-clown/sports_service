import Vue from 'vue'
import Router from 'vue-router'
import VideoPlayer from 'vue-video-player'
require('video.js/dist/video-js.css')
require('vue-video-player/src/custom-theme.css')

Vue.use(VideoPlayer)

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index2'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  // {
  //   path: '/',
  //   component: Layout,
  //   redirect: '/dashboard',
  //   children: [{
  //     path: 'dashboard',
  //     name: 'Dashboard',
  //     component: () => import('@/views/dashboard/index'),
  //     meta: { title: 'Dashboard', icon: 'dashboard' }
  //   }]
  // },
  //
  // {
  //   path: '/example',
  //   component: Layout,
  //   redirect: '/example/table',
  //   name: 'Example',
  //   meta: { title: 'Example', icon: 'el-icon-s-help' },
  //   children: [
  //     {
  //       path: 'table',
  //       name: 'Table',
  //       component: () => import('@/views/table/index'),
  //       meta: { title: 'Table', icon: 'table' }
  //     },
  //     {
  //       path: 'tree',
  //       name: 'Tree',
  //       component: () => import('@/views/tree/index'),
  //       meta: { title: 'Tree', icon: 'tree' }
  //     }
  //   ]
  // },
  //
  // {
  //   path: '/form',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'index',
  //       name: 'Form',
  //       component: () => import('@/views/form/index'),
  //       meta: { title: 'Form', icon: 'form' }
  //     }
  //   ]
  // },
  // {
  //   path: '/nested',
  //   component: Layout,
  //   redirect: '/nested/menu1',
  //   name: 'Nested',
  //   meta: {
  //     title: 'Nested',
  //     icon: 'nested'
  //   },
  //   children: [
  //     {
  //       path: 'menu1',
  //       component: () => import('@/views/nested/menu1/index'), // Parent router-view
  //       name: 'Menu1',
  //       meta: { title: 'Menu1' },
  //       children: [
  //         {
  //           path: 'menu1-1',
  //           component: () => import('@/views/nested/menu1/menu1-1'),
  //           name: 'Menu1-1',
  //           meta: { title: 'Menu1-1' }
  //         },
  //         {
  //           path: 'menu1-2',
  //           component: () => import('@/views/nested/menu1/menu1-2'),
  //           name: 'Menu1-2',
  //           meta: { title: 'Menu1-2' },
  //           children: [
  //             {
  //               path: 'menu1-2-1',
  //               component: () => import('@/views/nested/menu1/menu1-2/menu1-2-1'),
  //               name: 'Menu1-2-1',
  //               meta: { title: 'Menu1-2-1' }
  //             },
  //             {
  //               path: 'menu1-2-2',
  //               component: () => import('@/views/nested/menu1/menu1-2/menu1-2-2'),
  //               name: 'Menu1-2-2',
  //               meta: { title: 'Menu1-2-2' }
  //             }
  //           ]
  //         },
  //         {
  //           path: 'menu1-3',
  //           component: () => import('@/views/nested/menu1/menu1-3'),
  //           name: 'Menu1-3',
  //           meta: { title: 'Menu1-3' }
  //         }
  //       ]
  //     },
  //     {
  //       path: 'menu2',
  //       component: () => import('@/views/nested/menu2/index'),
  //       name: 'Menu2',
  //       meta: { title: 'menu2' }
  //     }
  //   ]
  // },

  {
    path: '/video',
    component: Layout,
    redirect: '/video',
    name: '视频',
    meta: {
      title: '视频模块',
      icon: 'eye-open'
    },
    children: [
      {
        path: 'list',
        component: () => import('@/views/video/index'),
        name: '视频管理',
        meta: { title: '视频管理', icon: 'list', affix: true }
      },
      {
        path: 'review',
        component: () => import('@/views/video/review'),
        name: '视频审核',
        meta: { title: '视频审核', icon: 'edit', affix: true }
      },
      {
        path: 'label',
        component: () => import('@/views/video/label'),
        name: '视频标签',
        meta: { title: '视频标签', icon: 'star', affix: true }
      }
    ]
  },

  {
    path: '/comment',
    component: Layout,
    redirect: '/comment',
    name: '评论',
    meta: {
      title: '评论模块',
      icon: 'message'
    },
    children: [
      {
        path: 'list',
        component: () => import('@/views/comment/comment'),
        name: '视频评论',
        meta: { title: '视频评论', icon: 'message', affix: true }
      },
      {
        path: 'barrage',
        component: () => import('@/views/comment/barrage'),
        name: '视频弹幕',
        meta: { title: '视频弹幕', icon: 'eye', affix: true }
      }
    ]
  },

  {
    path: '/user',
    component: Layout,
    redirect: '/user',
    name: '用户',
    meta: {
      title: '用户模块',
      icon: 'peoples'
    },
    children: [
      {
        path: 'list',
        component: () => import('@/views/user/user'),
        name: '用户管理',
        meta: { title: '用户管理', icon: 'people', affix: true }
      }
    ]
  },
  {
    path: '/configure',
    component: Layout,
    redirect: '/configure',
    name: '设置',
    meta: {
      title: '设置',
      icon: 'skill'
    },
    children: [
      {
        path: 'list',
        component: () => import('@/views/configure/hotSearch'),
        name: '热搜',
        meta: { title: '热搜', icon: 'search', affix: true }
      },
      {
        path: 'banner',
        component: () => import('@/views/configure/banner'),
        name: 'banner',
        meta: { title: 'banner', icon: 'eye-open', affix: true }
      },
      {
        path: 'avatar',
        component: () => import('@/views/configure/avatar'),
        name: '系统头像',
        meta: { title: '系统头像', icon: 'peoples', affix: true }
      }
    ],
  },

  // {
  //   path: 'external-link',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
  //       meta: { title: 'External Link', icon: 'link' }
  //     }
  //   ]
  // },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
