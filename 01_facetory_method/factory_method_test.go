package facetory_method

import "testing"

func TestGetGun(t *testing.T) {
	type fields struct {
		GunType string
	}
	type want struct {
		Name  string
		Power int
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "get ak47",
			fields: fields{
				GunType: "ak47",
			},
			want: want{
				Name:  "Ak47 gun",
				Power: 4,
			},
		},
		{
			name: "get m4a1",
			fields: fields{
				GunType: "m4a1",
			},
			want: want{
				Name:  "M4a1 gun",
				Power: 1,
			},
		},
		{
			name: "get unknown",
			fields: fields{
				GunType: "grenade",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGun(tt.fields.GunType)
			if err != nil {
				if err.Error() != "wrong gun type" {
					t.Errorf("getGun() = %v, want %v", err, "wrong gun type")
				}
				return
			}
			if got.getName() != tt.want.Name {
				t.Errorf("getName() = %v, want %v", got.getName(), tt.want.Name)
			}
			if got.getPower() != tt.want.Power {
				t.Errorf("getPower() = %v, want %v", got.getPower(), tt.want.Power)
			}
		})
	}
}
