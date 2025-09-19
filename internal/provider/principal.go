// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import "github.com/hashicorp/boundary/api/roles"

func flattenRolePrincipalsKeyInfo(principals []*roles.Principal) []interface{} {
	if principals == nil || len(principals) == 0 {
		return []interface{}{}
	}

	result := make([]interface{}, 0, len(principals))
	for _, p := range principals {
		if p == nil {
			continue
		}
		m := make(map[string]interface{})

		if v := p.Id; v != "" {
			m[IDKey] = v
		}
		if v := p.Type; v != "" {
			m[TypeKey] = v
		}
		if v := p.ScopeId; v != "" {
			m[ScopeIdKey] = v
		}

		result = append(result, m)
	}

	return result
}
