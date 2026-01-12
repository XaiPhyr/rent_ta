import type { Component } from 'vue';

interface RouteInterface {
  path: string;
  name: string;
  title: string;
  icon?: string;
  component?: Component;
  callback?: () => void;
  beforeEnter?: (to: any, from: any, next: any) => void;
  meta?: {
    layout: Component;
  };
}

export interface RoutesInterface extends Array<RouteInterface> {}
