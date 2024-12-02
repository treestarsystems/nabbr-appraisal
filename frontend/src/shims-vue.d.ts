declare module '*.vue' {
  import { DefineComponent } from 'vue';
  declare module 'postscribe';
  const component: DefineComponent<{}, {}, any>;
  export default component;
}
