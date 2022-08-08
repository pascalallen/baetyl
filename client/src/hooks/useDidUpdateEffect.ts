import { DependencyList, EffectCallback, useEffect, useRef } from 'react';

type Destructor = () => void;

const useDidUpdateEffect = (effect: EffectCallback, deps?: DependencyList): void => {
  const didMountRef = useRef(false);

  useEffect((): void | Destructor => {
    if (!didMountRef.current) {
      didMountRef.current = true;
    } else {
      return effect();
    }
  }, [effect, deps]);
};

export default useDidUpdateEffect;
