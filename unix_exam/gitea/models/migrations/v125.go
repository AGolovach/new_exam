// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"xorm.io/xorm"
)

func addReviewMigrateInfo(x *xorm.Engine) error {
	type Review struct {
		OriginalAuthor   string
		OriginalAuthorID int64
	}

	if err := x.Sync2(new(Review)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}
	return nil
}
