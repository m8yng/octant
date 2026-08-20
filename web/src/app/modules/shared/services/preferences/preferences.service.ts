/*
 * Copyright (c) 2020 the Octant contributors. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

import { Injectable, OnDestroy } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { PreferencePanel, Preferences } from '../../models/preference';
import { ThemeService } from '../theme/theme.service';
import { PreferencesEntry } from './preferences.entry';

@Injectable({
  providedIn: 'root',
})
export class PreferencesService implements OnDestroy {
  private kubeConfigPathText: string;
  private kubeConfigFullPath: string;

  public preferences: Map<string, PreferencesEntry<any>> = new Map();
  public preferencesOpened: BehaviorSubject<boolean> =
    new BehaviorSubject<boolean>(false);

  constructor(private themeService: ThemeService) {
    this.preferences.set(
      'navigation.collapsed',
      new PreferencesEntry<boolean>(
        this,
        'navigation.collapsed',
        false,
        'collapsed'
      )
    );

    this.preferences.set(
      'navigation.labels',
      new PreferencesEntry<boolean>(this, 'navigation.labels', true, 'show')
    );

    this.preferences.set(
      'general.pageSize',
      new PreferencesEntry<number>(this, 'general.pageSize', 10, '')
    );

    this.preferences.set(
      'theme',
      new PreferencesEntry<string>(
        this,
        'theme',
        this.themeService.themeType.value,
        ''
      )
    );

    this.updateTheme();
  }

  ngOnDestroy(): void {
    for (const pref of this.preferences.values()) {
      pref.destroy();
    }
  }

  setStoredValue(key: string, value: any) {
    localStorage.setItem(key, value);
  }

  getStoredValue(key: string, defaultValue: any) {
    return localStorage.getItem(key) || defaultValue;
  }

  public preferencesChanged(update: Preferences) {
    for (const pref of this.preferences.values()) {
      pref.preferencesChanged(update);
    }
    this.updateTheme();
  }

  public getPreferences(): Preferences {
    return {
      updateName: 'generalPreferences',
      panels: [this.getGeneralPanels(), this.getNavigationPanels()],
    };
  }

  public setKubeConfigPath(path: string) {
    this.kubeConfigFullPath = path;
    this.kubeConfigPathText = path;
    if (path.length > 35) {
      this.kubeConfigPathText = `${path.substring(0, 17)}...${path.substring(
        path.length - 17,
        path.length
      )}`;
    }
  }

  updateTheme() {
    if (
      this.themeService.themeType.value !==
      this.preferences.get('theme').subject.value
    ) {
      this.themeService.switchTheme();
    }
  }

  reset() {
    for (const pref of this.preferences.values()) {
      pref.setDefaultValue();
    }
    this.updateTheme();
  }

  private getGeneralPanels(): PreferencePanel {
    const pageSize = this.preferences.get('general.pageSize').subject.value;

    return {
      name: 'General',
      sections: [
        {
          name: 'Color Theme',
          elements: [
            {
              name: 'theme',
              type: 'radio',
              value: this.preferences.get('theme').subject.value,
              config: {
                values: [
                  {
                    label: 'Dark',
                    value: 'dark',
                  },
                  {
                    label: 'Light',
                    value: 'light',
                  },
                ],
              },
            },
          ],
        },
        {
          name: 'Table Pagination',
          elements: [
            {
              name: 'general.pageSize',
              value: pageSize,
              label: 'Page size:',
              type: 'dropdown',
              metadata: {
                type: 'dropdown',
                title: [
                  {
                    metadata: {
                      type: 'input',
                    },
                    config: {
                      value: pageSize,
                    },
                  },
                ],
              },
              config: {
                type: 'label',
                selection: pageSize,
                useSelection: true,
                items: [
                  {
                    name: '10',
                    type: 'text',
                    label: '10',
                  },
                  {
                    name: '20',
                    type: 'text',
                    label: '20',
                  },
                  {
                    name: '50',
                    type: 'text',
                    label: '50',
                  },
                  {
                    name: '100',
                    type: 'text',
                    label: '100',
                  },
                ],
              },
            },
          ],
        },
        {
          name: 'Current Kube Config Path',
          elements: [
            {
              type: 'text',
              name: 'general.kubeConfig',
              value: '',
              textConfig: {
                config: {
                  value: this.kubeConfigPathText || 'unknown',
                  clipboardValue: this.kubeConfigFullPath || 'unknown',
                },
              },
            },
          ],
        },
      ],
    };
  }

  private getNavigationPanels(): PreferencePanel {
    return {
      name: 'Navigation',
      sections: [
        {
          name: 'Navigation',
          elements: [
            {
              name: 'navigation.collapsed',
              type: 'radio',
              value: this.preferences.get('navigation.collapsed').subject.value
                ? 'collapsed'
                : 'expanded',
              config: {
                values: [
                  {
                    label: 'Expanded',
                    value: 'expanded',
                  },
                  {
                    label: 'Collapsed',
                    value: 'collapsed',
                  },
                ],
              },
            },
          ],
        },
        {
          name: 'Navigation labels',
          elements: [
            {
              name: 'navigation.labels',
              type: 'radio',
              value: this.preferences.get('navigation.labels').subject.value
                ? 'show'
                : 'hide',
              config: {
                values: [
                  {
                    label: 'Show Labels',
                    value: 'show',
                  },
                  {
                    label: 'Hide Labels',
                    value: 'hide',
                  },
                ],
              },
            },
          ],
        },
      ],
    };
  }
}
