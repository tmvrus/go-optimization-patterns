package zerocopy

import (
	"io"
)

type badReader struct {
	offset int
	data   []byte
}

func (r *badReader) Read() []byte {
	dst := make([]byte, 0, 100)
	n := copy(dst, r.data[r.offset:len(dst)])
	r.offset = r.offset + n
	if n == 0 {
		return nil
	}
	return dst
}

func newBadReader() *badReader {
	return &badReader{
		offset: 0,
		data:   dataSource,
	}
}

type goodReader struct {
	offset int
	data   []byte
}

func (r *goodReader) Read(dst []byte) error {
	n := copy(dst, r.data[r.offset:len(dst)])
	r.offset = r.offset + n
	if n == 0 {
		return io.EOF
	}
	return nil
}

func newGoodReader() *goodReader {
	return &goodReader{
		offset: 0,
		data:   dataSource,
	}
}

var dataSource = []byte(`QuotaAmount,StartDate,OwnerName,Username
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
QuotaAmount,StartDate,OwnerName,Username
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
QuotaAmount,StartDate,OwnerName,Username
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
QuotaAmount,StartDate,OwnerName,Username
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
QuotaAmount,StartDate,OwnerName,Username
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-01-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-02-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-03-01,Chris Riley,trailhead9.ub20k5i9t8ou@example.com
150000,2016-01-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-02-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-03-01,Harold Campbell,trailhead14.jibpbwvuy67t@example.com
150000,2016-01-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-02-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-03-01,Jessica Nichols,trailhead19.d1fxj2goytkp@example.com
150000,2016-01-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-02-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-03-01,Catherine Brown,trailhead16.kojyepokybge@example.com
150000,2016-01-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-02-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-03-01,Kelly Frazier,trailhead7.zdcsy4ax10mr@example.com
150000,2016-01-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-02-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com
150000,2016-03-01,Dennis Howard,trailhead4.wfokpckfroxp@example.com`)
