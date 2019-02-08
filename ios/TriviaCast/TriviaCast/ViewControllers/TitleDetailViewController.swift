//
//  TitleDetailViewController.swift
//  TriviaCast
//
//  Created by Brophy on 2/7/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit
import SnapKit

class TitleDetailViewController: TCViewController {
    
    private let titleLabel: UILabel = UILabel()
    private let detailStackView: UIStackView = UIStackView()
    
    private var detailLabels: [UILabel] = []
    
    override func viewDidLoad() {
        super.viewDidLoad()
        build()
        configure()
    }
    
    func setTitle(_ title: String?) {
        titleLabel.text = title
    }
    
    func setBulletPoints(_ texts: [String]) {
        for label in detailLabels {
            label.removeFromSuperview()
        }
        detailLabels = []
        for text in texts {
            let label = UILabel()
            label.font = TCFont.h3
            label.text = "- \(text)"
            detailLabels.append(label)
            detailStackView.addArrangedSubview(label)
        }
    }
    
    // MARK: - Private Interface
    
    private func build() {
        view.addSubview(titleLabel)
        view.addSubview(detailStackView)
        
        titleLabel.font = TCFont.h1
        titleLabel.numberOfLines = 0
        titleLabel.textAlignment = .center
        
        detailStackView.axis = .vertical
        detailStackView.alignment = .leading
        detailStackView.spacing = SpacingUnit.u1
    }
    
    private func configure() {
        titleLabel.snp.makeConstraints { make in
            make.leading.trailing.equalTo(view)
            make.top.equalTo(view).offset(SpacingUnit.u2)
        }
        
        detailStackView.snp.makeConstraints { make in
            make.top.equalTo(titleLabel.snp.bottom).offset(SpacingUnit.u2)
            make.leading.equalTo(view).offset(SpacingUnit.u2)
            make.trailing.bottom.equalTo(view).offset(-SpacingUnit.u2)
        }
    }
}
