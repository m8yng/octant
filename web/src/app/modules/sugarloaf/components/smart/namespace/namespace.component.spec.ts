// Copyright (c) 2019 the Octant contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { NamespaceComponent } from './namespace.component';
import { NgSelectModule } from '@ng-select/ng-select';
import { windowProvider, WindowToken } from '../../../../../window';
import { SharedModule } from 'src/app/modules/shared/shared.module';
import {
  OverlayScrollbarsComponent,
  OverlayscrollbarsModule,
} from 'overlayscrollbars-ngx';

describe('NamespaceComponent', () => {
  let component: NamespaceComponent;
  let fixture: ComponentFixture<NamespaceComponent>;

  beforeEach(
    waitForAsync(() => {
      TestBed.configureTestingModule({
        imports: [NgSelectModule, OverlayscrollbarsModule],
        declarations: [NamespaceComponent, OverlayScrollbarsComponent],
        providers: [{ provide: WindowToken, useFactory: windowProvider }],
      }).compileComponents();
    })
  );

  beforeEach(() => {
    fixture = TestBed.createComponent(NamespaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('shows all namespaces for an all-namespaces route', () => {
    component.activeUrl = '/overview/all-namespaces/workloads/pods';

    expect(component.namespaceLabel()).toBe('All Namespaces');
    expect(component.namespaceClass('', true)).toContain('active');
  });
});
