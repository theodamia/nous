import { useLayoutEffect } from 'react';

/**
 * Hook to prevent scroll locking when drawer/dialog is open
 *
 * PROBLEM:
 * shadcn's Drawer component (built on Vaul + Radix Dialog) uses react-remove-scroll
 * to lock body scrolling when a drawer opens. This involves:
 * 1. Setting body styles: overflow: hidden, position: relative, margin-right (for scrollbar)
 * 2. Adding JavaScript event listeners to prevent wheel/touch events
 * 3. Adding data-scroll-locked="1" attribute to body element to indicate scroll is locked
 *
 * REQUIREMENT:
 * We need background scrolling to remain enabled for low-resolution screens where
 * users need to scroll the page content behind the drawer, while keeping the overlay visible.
 *
 * WHY THE HACK IS NEEDED:
 * - CSS alone (overflow: auto !important) fixes the scrollbar visibility but doesn't
 *   re-enable trackpad/wheel scrolling because react-remove-scroll's JavaScript event
 *   listeners call preventDefault() on wheel/touchmove events
 * - Vaul's disablePreventScroll prop only affects input repositioning, not body scroll lock
 * - The scroll lock comes from Radix Dialog's internal use of react-remove-scroll,
 *   which doesn't expose a way to disable just the scroll lock while keeping modal behavior
 *
 * SOLUTION:
 * This hook intercepts scroll events in the capture phase (before react-remove-scroll's
 * listeners) and stops propagation when data-scroll-locked="1" is present, effectively
 * preventing react-remove-scroll from blocking the scroll while allowing the native
 * browser scroll behavior to continue.
 *
 * Combined with, this restores full scroll functionality
 * layer base {
 *  body[data-scroll-locked='1'] {
 *     overflow: auto !important;
 *   }
 * }
 * (both scrollbar and trackpad) while maintaining the drawer's modal overlay.
 */

const usePreventScrollLock = () => {
  useLayoutEffect(() => {
    const handler = (e: Event) => {
      if (document.body.getAttribute('data-scroll-locked') === '1') {
        e.stopImmediatePropagation();
      }
    };

    const eventTypes = ['wheel', 'scroll', 'touchmove'] as const;

    const options = { capture: true, passive: false };
    eventTypes.forEach((type) => {
      window.addEventListener(type, handler, options);
    });

    return () => {
      eventTypes.forEach((type) => {
        window.removeEventListener(type, handler, options);
      });
    };
  }, []);
};

export default usePreventScrollLock;
