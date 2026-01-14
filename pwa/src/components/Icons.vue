<script setup lang="ts">
  import MenuIcon from '@/svg/bar.svg';
  import BellIcon from '@/svg/bell.svg';
  import ArrowLeftIcon from '@/svg/arrow-left.svg';
  import MagnifyingGlassIcon from '@/svg/magnifying-glass.svg';
  import XIcon from '@/svg/x-mark.svg';
  import DollarIcon from '@/svg/dollar.svg';
  import ArrowsPointingOutIcon from '@/svg/arrows-pointing-out.svg';
  import ArrowsPointingInIcon from '@/svg/arrows-pointing-in.svg';
  import PinIcon from '@/svg/pin.svg';
  import AtIcon from '@/svg/at-symbol.svg';
  import ClockIcon from '@/svg/clock.svg';
  import HomeIcon from '@/svg/home.svg';
  import UserIcon from '@/svg/user.svg';
  import HeartIcon from '@/svg/heart.svg';
  import WalletIcon from '@/svg/wallet.svg';

  const props = defineProps({
    icon: { type: String, required: true, default: '' },
    alt: { type: String, default: '' },
    title: { type: String, default: '' },
  });

  const items: Record<string, any> = {
    menu: MenuIcon,
    bell: BellIcon,
    'arrow-left': ArrowLeftIcon,
    'magnifying-glass': MagnifyingGlassIcon,
    dollar: DollarIcon,
    close: XIcon,
    'arrows-pointing-out': ArrowsPointingOutIcon,
    'arrows-pointing-in': ArrowsPointingInIcon,
    email: AtIcon,
    pin: PinIcon,
    clock: ClockIcon,
    home: HomeIcon,
    user: UserIcon,
    heart: HeartIcon,
    wallet: WalletIcon,
  };
</script>

<template>
  <main>
    <img
      :src="items[props.icon]"
      :alt="props.alt"
      :title="props.title"
      class="w-5 h-5"
    />
  </main>
</template>
